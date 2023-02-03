package csv

import (
	"encoding/csv"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"io"
	"os"
	"strconv"
)

func ReadCSV(ctx *fiber.Ctx) error {
	f, err := os.Open("csv/csv1.csv")
	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			panic(err.Error())
		}
	}(f)

	var addresses []Address

	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == nil {
			pin, _ := strconv.Atoi(rec[5])
			addresses = append(
				addresses, Address{
					firstName: rec[0],
					lastName:  rec[1],
					line1:     rec[2],
					line2:     rec[3],
					city:      rec[4],
					pincode:   pin,
				})
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return ctx.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	response := "--------------------------------------------------------------------------------------------------------\n"
	response += "| First Name    | Last Name 	| Line 1 			| Line 2 			| City 				| PinCode 	|\n"

	for _, v := range addresses {
		response += fmt.Sprintf("| %s | %s | %s | %s | %s | %d |\n", v.firstName, v.lastName, v.line1, v.line2, v.city, v.pincode)
	}

	return ctx.Status(200).SendString(response)
}
