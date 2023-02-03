package handler

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
)

func init() {
	var err error
	bookDB, err = sql.Open("postgre", "")
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
}

var bookDB *sql.DB

func createBook(ctx *fiber.Ctx) error {

}

func getBook(ctx *fiber.Ctx) error {

}

func updateBook(ctx *fiber.Ctx) error {

}

func deleteBook(ctx *fiber.Ctx) error {

}
