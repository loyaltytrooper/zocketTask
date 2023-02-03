package routes

var (
	root = "/"

	createBook  = ""
	createMovie = ""
	createFood  = ""
	getBook     = ""
	getMovie    = ""
	getFood     = ""
	updateBook  = ""
	updateMovie = ""
	updateFood  = ""
	deleteBook  = ""
	deleteMovie = ""
	deleteFood  = ""

	readCSV = "/csv/read"

	getImagesBase64 = "/images/base64" // API that will get base64 encoded images from the URLs provided
	uploadFiles     = "/files/upload"  // body that will take an array of URLs
)
