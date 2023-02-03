package files

import (
	"github.com/gofiber/fiber/v2"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type ImageURLs struct {
	URLs []string `json:"urls" form:"urls"`
}

var downloadChannel = make(chan string, 15)

func DownloadFiles(ctx *fiber.Ctx) error {
	var urls ImageURLs
	if err := ctx.BodyParser(&urls); err != nil {
		log.Fatalf("Error parsing the payload request: %s", err.Error())
		return ctx.SendStatus(422)
	}
	log.Println("---1")
	for _, v := range urls.URLs {
		log.Println("---3")
		go downloadFile(v)
		log.Println("---4")
	}
	log.Println("---2")
	var images []string
	for i := 0; i < len(urls.URLs); i++ {
		log.Println("---5")
		x := <-downloadChannel
		log.Println("---6")
		images = append(images, x)
		log.Println(images)
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
