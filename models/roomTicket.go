package models

import (
	"time"
)

type RoomTicket struct {
	ID           int
	Number       int
	Avaliability AvaliabilityStatus
	Price        float32
	TotalAmount  float32
	CheckInTime  time.Time
	CheckOutTime time.Time
}
