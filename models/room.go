package models

type Room struct {
	ID           int
	Number       int
	Avaliability AvaliabilityStatus
	Price        float32
	Period       int
}
