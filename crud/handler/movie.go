package handler

import (
	"github.com/gofiber/fiber/v2"
	"zocketTask/crud/database"
	"zocketTask/crud/models"
)

func CreateMovie(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"status": "Hello"})
}

func GetMovie(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"status": "Hello"})
}

func GetAllMovies(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	bookPayload := new(models.Book)

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "select * from book where id=$1", bookID).Scan(&bookPayload.Title, &bookPayload.Type, &bookPayload.Author, &bookPayload.Price, &bookPayload.ID)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(bookPayload)
}

func UpdateMovie(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"status": "Hello"})
}

func DeleteMovie(ctx *fiber.Ctx) error {

	return ctx.JSON(fiber.Map{"status": "Hello"})
}
