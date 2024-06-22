package models

type RoomTicket struct {
	ID           int
	Number       int
	Avaliability AvaliabilityStatus
	Price        float32
	TotalAmount  float64
	Period       int
}
