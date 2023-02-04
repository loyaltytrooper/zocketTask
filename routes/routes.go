package routes

import (
	"github.com/gofiber/fiber/v2"
	"zocketTask/crud/handler"
	"zocketTask/csv"
	"zocketTask/files"
	rootFunc "zocketTask/root"
)

func RegisterRoutes(app *fiber.App) {
	app.Get(root, rootFunc.Hello)

	app.Post(book, handler.CreateBook)
	app.Post(movie, handler.CreateMovie)

	app.Post(books, handler.CreateMoreBooks)
	app.Post(movies, handler.CreateMoreMovies)

	app.Get(bookID, handler.GetBook)
	app.Get(movieID, handler.GetMovie)

	app.Get(book, handler.GetAllBooks)
	app.Get(movie, handler.GetAllMovies)

	app.Put(book, handler.UpdateBook)
	app.Put(movie, handler.UpdateMovie)

	app.Delete(bookID, handler.DeleteBook)
	app.Delete(movieID, handler.DeleteMovie)

	app.Get(readCSV, csv.ReadCSV)

	app.Post(getImagesBase64, files.DownloadFiles)
	app.Get(getImagesBase64, files.GetImage)
}
