package files

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http"
	"os"
	"path"
)

var downloadChannel = make(chan string, 15)

func GetImage(ctx *fiber.Ctx) error {
	var file URL
	if err := ctx.BodyParser(&file); err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if _, err := os.Stat(file.File); err == nil {
		return ctx.Status(200).SendFile(file.File)
	} else {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "File not found, try re-downloading",
		})
	}
}

func DownloadFiles(ctx *fiber.Ctx) error {
	var urls ImageURLs
	if err := ctx.BodyParser(&urls); err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	for _, v := range urls.URLs {
		go downloadFile(v)
	}
	var images []string
	for i := 0; i < len(urls.URLs); i++ {
		x := <-downloadChannel
		images = append(images, x)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "Success",
		"images": images,
	})
}

func downloadFile(url string) {
	resp, err := http.Get(url)
	if err != nil {
		downloadChannel <- "failed"
		return
	}
	defer resp.Body.Close()

	// Create the file
	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)
	if err != nil {
		downloadChannel <- "failed"
		return
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	downloadChannel <- filepath
	return
}
