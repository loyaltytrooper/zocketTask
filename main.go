package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"zocketTask/routes"
)

func main() {
	app := fiber.New()

	routes.RegisterRoutes(app)
	log.Fatal(app.Listen("127.0.0.1:9010"))
}
