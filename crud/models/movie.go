package models

type Movie struct {
	ID          int      `json:"id"`
	Category    string   `json:"category"`
	Name        string   `json:"name"`
	Actors      []string `json:"actors"` // or normal string but then cleaning operations have to be done
	TicketPrice float32  `json:"ticketPrice"`
}
