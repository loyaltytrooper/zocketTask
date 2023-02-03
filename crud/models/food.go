package models

type Food struct {
	ID      int     `json:"id"`
	Cuisine string  `json:"cuisine"`
	Name    string  `json:"name"`
	Price   float32 `json:"price"`
}
