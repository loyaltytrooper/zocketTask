package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"zocketTask/crud/database"
	"zocketTask/routes"
)

func main() {
	godotenv.Load()

	database.InitDB()
	defer database.CloseDB()

	app := fiber.New()

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen("127.0.0.1:9010"))
}
