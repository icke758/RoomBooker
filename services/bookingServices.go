package services

import (
	"room/database"
	"room/models"
)

func BookRoom(roomNumber int, period int) bool {
	booking := database.CheckBooking(roomNumber)
	if models.AvaliabilityStatusFromInt(booking) == models.Rented {
		return false
	}
	database.BookRoom(roomNumber, period)
	return true
}

func UnbookRoom(roomNumber int) bool {
	booking := database.CheckBooking(roomNumber)
	if models.AvaliabilityStatusFromInt(booking) == models.Available {
		return false
	}
	database.UnbookRoom(roomNumber)
	return true
}
