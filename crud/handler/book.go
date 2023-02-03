package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"zocketTask/crud/database"
	"zocketTask/crud/models"
)

func CreateBook(ctx *fiber.Ctx) error {
	bookPayload := new(models.Book)
	if err := ctx.BodyParser(&bookPayload); err != nil {
		log.Fatalf("Error parsing the payload request: %s", err.Error())
		return ctx.SendStatus(422)
	}

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "insert into book(title, author, type, price)  values ($1, $2, $3, $4) returning (id)", bookPayload.Title, bookPayload.Author, bookPayload.Type, bookPayload.Price).Scan(&bookPayload.ID)
	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(bookPayload)
}

func CreateMoreBooks(ctx *fiber.Ctx) error {
	var booksPayload models.Books
	if err := ctx.BodyParser(&booksPayload); err != nil {
		log.Fatalf("Error parsing the payload request: %s", err.Error())
		return ctx.SendStatus(422)
	}

	query := fmt.Sprintf("insert into book(title, author, type, price) ")
	if len(booksPayload.Books) > 0 {
		query += fmt.Sprintf("values ('%s', '%s', '%s', %f)", booksPayload.Books[0].Title, booksPayload.Books[0].Author, booksPayload.Books[0].Type, booksPayload.Books[0].Price)
		for i := 1; i < len(booksPayload.Books); i++ {
			query += fmt.Sprintf(", ('%s', '%s', '%s', %f)", booksPayload.Books[i].Title, booksPayload.Books[i].Author, booksPayload.Books[i].Type, booksPayload.Books[i].Price)
		}
	}

	log.Print(query)
	_, err := database.CRUD_DB.Query(ctx.UserContext(), query)
	log.Print("----")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": (err.Error()),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "Success",
	})
}

func GetBook(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	bookPayload := new(models.Book)

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "select * from book where id=$1", bookID).Scan(&bookPayload.Title, &bookPayload.Type, &bookPayload.Author, &bookPayload.Price, &bookPayload.ID)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(200).JSON(bookPayload)
}

func GetAllBooks(ctx *fiber.Ctx) error {
	rows, err := database.CRUD_DB.Query(ctx.UserContext(), "select * from book")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	var bookPayload []models.Book

	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.Title, &book.Type, &book.Author, &book.Price, &book.ID); err != nil {
			log.Fatal(err)
		}
		bookPayload = append(bookPayload, book)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"books": bookPayload,
	})
}

func UpdateBook(ctx *fiber.Ctx) error {
	bookPayload := new(models.Book)
	if err := ctx.BodyParser(&bookPayload); err != nil {
		log.Fatalf("Error parsing the payload request: %s", err.Error())
		return ctx.SendStatus(422)
	}

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "update book set title=$1, author=$2, type=$3, price=$4 WHERE id = $5 returning *", bookPayload.Title, bookPayload.Author, bookPayload.Type, bookPayload.Price, bookPayload.ID).Scan(&bookPayload.Title, &bookPayload.Author, &bookPayload.Type, &bookPayload.Price, &bookPayload.ID)
	if err != nil {
		log.Fatalf("Error executing the update query: %s", err.Error())
		return ctx.SendStatus(422)
	}

	return ctx.Status(200).JSON(bookPayload)
}

func DeleteBook(ctx *fiber.Ctx) error {
	bookID := ctx.Params("id")
	bookPayload := new(models.Book)

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "delete from book where id=$1 returning *", bookID).Scan(&bookPayload.Title, &bookPayload.Type, &bookPayload.Author, &bookPayload.Price, &bookPayload.ID)
	if err != nil {
		log.Print(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(bookPayload)
}
