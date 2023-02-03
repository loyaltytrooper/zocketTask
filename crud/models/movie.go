package models

type Movie struct {
	ID          int      `json:"id" form:"id"`
	Name        string   `json:"name" form:"name"`
	Director    string   `json:"director" form:"director"`
	Type        string   `json:"category" form:"category"`
	Actors      []string `json:"actors" form:"actors"` // or normal string but then cleaning operations have to be done
	TicketPrice float32  `json:"ticket_price" form:"ticket_price"`
}

type Movies struct {
	Movies []Movie `json:"movies" form:"movies"`
}
