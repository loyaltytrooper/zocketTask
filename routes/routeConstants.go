package routes

var (
	root = "/"

	book  = "/book"
	books = "/books"

	movie  = "/movie"
	movies = "/movies"

	bookID  = "/book/:id"
	movieID = "/movie/:id"

	readCSV = "/csv/read"

	getImagesBase64 = "/images/base64" // API that will return base64 encoded images from the URLs provided

	// will implement this as my own small side project after 22 Feb because of end semesters
	uploadFiles = "/files/upload" // body that will take an array of URLs
)
