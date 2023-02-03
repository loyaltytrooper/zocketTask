package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"time"
	"zocketTask/crud/database"
	"zocketTask/routes"
)

func main() {
	godotenv.Load()

	database.InitDB()
	defer database.CloseDB()

	app := fiber.New()
	err := app.ShutdownWithTimeout(time.Duration(time.Second * 10))
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen("127.0.0.1:9010"))
}
