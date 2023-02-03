package models

type Book struct {
	ID       int     `json:"id"`
	Category string  `json:"category"`
	Name     string  `json:"name"`
	Author   string  `json:"author"`
	Price    float32 `json:"price"`
}
