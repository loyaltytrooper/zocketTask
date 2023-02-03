package main

import "github.com/gofiber/fiber/v2"

func Error(err error) fiber.Map {
	return fiber.Map{
		"status": "Not Working",
		"error":  err.Error(),
	}
}
