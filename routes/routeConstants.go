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

	getImagesBase64 = "/images/base64" // API that will get base64 encoded images from the URLs provided
	uploadFiles     = "/files/upload"  // body that will take an array of URLs
)
