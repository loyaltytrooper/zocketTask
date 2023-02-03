package routes

import (
	"github.com/gofiber/fiber/v2"
	rootFunc "zocketTask/root"
)

func RegisterRoutes(app *fiber.App) {
	app.Get(root, rootFunc.Hello)
}
