package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"zocketTask/crud/database"
	"zocketTask/crud/models"
)

func CreateMovie(ctx *fiber.Ctx) error {
	moviePayload := new(models.Movie)
	if err := ctx.BodyParser(&moviePayload); err != nil {
		return ctx.Status(422).JSON(fiber.Map{
			"error": fmt.Sprintf("Error parsing the payload request: %s", err.Error()),
		})
	}

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "insert into movie(name, director, type, actors, ticketprice)  values ($1, $2, $3, $4) returning (id)", moviePayload.Name, moviePayload.Director, moviePayload.Type, moviePayload.Actors, moviePayload.TicketPrice).Scan(&moviePayload.ID)
	if err != nil {
		return ctx.Status(4000).JSON(fiber.Map{
			"error": fmt.Sprintf("Error executing Database Query: %s", err.Error()),
		})
	}

	return ctx.Status(200).JSON(moviePayload)
}

func CreateMoreMovies(ctx *fiber.Ctx) error {
	var moviesPayload models.Movies
	if err := ctx.BodyParser(&moviesPayload); err != nil {
		log.Fatalf("Error parsing the payload request: %s", err.Error())
		return ctx.SendStatus(422)
	}

	query := fmt.Sprint("insert into movie(name, director, type, ticketprice, actors) ")

	if len(moviesPayload.Movies) > 0 {
		query += fmt.Sprintf("values ('%s', '%s', '%s', %f, ARRAY [", moviesPayload.Movies[0].Name, moviesPayload.Movies[0].Director, moviesPayload.Movies[0].Type, moviesPayload.Movies[0].TicketPrice)
		query += "'" + moviesPayload.Movies[0].Actors[0] + "'"
		for x := 1; x < len(moviesPayload.Movies[0].Actors); x++ {
			query += ", '" + moviesPayload.Movies[0].Actors[x] + "'"
		}
		query += "])"
		for i := 1; i < len(moviesPayload.Movies); i++ {
			query += fmt.Sprintf(", ('%s', '%s', '%s', %f, ARRAY [", moviesPayload.Movies[i].Name, moviesPayload.Movies[i].Director, moviesPayload.Movies[i].Type, moviesPayload.Movies[i].TicketPrice)

			query += "'" + moviesPayload.Movies[i].Actors[0] + "'"
			for j := 1; j < len(moviesPayload.Movies[i].Actors); j++ {
				query += ", '" + moviesPayload.Movies[i].Actors[j] + "'"
			}
			query += "])"
		}
	}

	log.Print(query)
	_, err := database.CRUD_DB.Query(ctx.UserContext(), query)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(fiber.Map{
		"status": "Success",
	})
}

func GetMovie(ctx *fiber.Ctx) error {
	movieID := ctx.Params("id")
	moviePayload := new(models.Movie)

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "select * from movie where id=$1", movieID).Scan(&moviePayload.Name, &moviePayload.Director, &moviePayload.Type, &moviePayload.Actors, &moviePayload.TicketPrice, &moviePayload.ID)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	return ctx.Status(200).JSON(moviePayload)
}

func GetAllMovies(ctx *fiber.Ctx) error {
	rows, err := database.CRUD_DB.Query(ctx.UserContext(), "select * from movie")
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var moviesPayload []models.Movie

	for rows.Next() {
		var movie models.Movie
		if err := rows.Scan(&movie.Name, &movie.Director, &movie.Type, &movie.Actors, &movie.TicketPrice, &movie.ID); err != nil {
			log.Fatal(err)
		}
		moviesPayload = append(moviesPayload, movie)
	}

	return ctx.Status(200).JSON(fiber.Map{
		"movies": moviesPayload,
	})
}

func UpdateMovie(ctx *fiber.Ctx) error {
	moviePayload := new(models.Movie)
	if err := ctx.BodyParser(&moviePayload); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "update movie set name=$1, director=$2, type=$3, actors=$4, ticketprice=$5 WHERE id = $5 returning *", moviePayload.Name, moviePayload.Director, moviePayload.Type, moviePayload.Actors, moviePayload.TicketPrice, moviePayload.ID).Scan(&moviePayload.Name, &moviePayload.Director, &moviePayload.Type, &moviePayload.Actors, &moviePayload.TicketPrice, &moviePayload.ID)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(moviePayload)
}

func DeleteMovie(ctx *fiber.Ctx) error {
	movieID := ctx.Params("id")
	movie := new(models.Movie)

	err := database.CRUD_DB.QueryRow(ctx.UserContext(), "delete from movie where id=$1 returning *", movieID).Scan(&movie.Name, &movie.Director, &movie.Type, &movie.Actors, &movie.TicketPrice, &movie.ID)
	if err != nil {
		log.Print(err.Error())
		return ctx.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(200).JSON(movie)
}
