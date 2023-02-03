package models

type Book struct {
	ID     int     `json:"id"  form:"id" `
	Type   string  `json:"category" form:"category"`
	Title  string  `json:"title" form:"title"`
	Author string  `json:"author" form:"author"`
	Price  float32 `json:"price" form:"price"`
}

type Books struct {
	Books []Book `json:"books" form:"books"`
}
