package models

type Movie struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Director    string   `json:"director"`
	Type        string   `json:"category"`
	Actors      []string `json:"actors"` // or normal string but then cleaning operations have to be done
	TicketPrice float32  `json:"ticket_price"`
}
