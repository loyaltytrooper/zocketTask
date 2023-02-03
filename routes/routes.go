package routes

import (
	"github.com/gofiber/fiber/v2"
	"zocketTask/crud/handler"
	rootFunc "zocketTask/root"
)

func RegisterRoutes(app *fiber.App) {
	app.Get(root, rootFunc.Hello)

	app.Post(book, handler.CreateBook)
	app.Post(movie, handler.CreateMovie)

	app.Post(books, handler.CreateMoreBooks)
	app.Post(movies, handler.CreateMoreBooks)

	app.Get(bookID, handler.GetBook)
	app.Get(movieID, handler.GetMovie)

	app.Get(book, handler.GetAllBooks)
	app.Get(movie, handler.GetAllMovies)

	app.Put(book, handler.UpdateBook)
	app.Put(movie, handler.UpdateMovie)

	app.Delete(bookID, handler.DeleteBook)
	app.Delete(movieID, handler.DeleteMovie)
}
