package services

import (
	"log"
	"room/database"
	"room/models"
)

func CreateReceipt(roomNumber int) models.RoomTicket {
	room, err := database.FindRoomByNumber(roomNumber)
	if err != nil {
		log.Fatal(err)
	}
	totalAmount := EvaluateStay(room.Number)

	receipt := models.RoomTicket{
		Number:       room.Number,
		Avaliability: room.Avaliability,
		Price:        room.Price,
		Period:       room.Period,
		TotalAmount:  totalAmount,
	}
	return receipt
}
