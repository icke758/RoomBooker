package services

import (
	"fmt"
	"room/database"
	"room/models"
)

func BookRoom(roomNumber int) bool {
	booking := database.CheckBooking(roomNumber)
	if models.AvaliabilityStatusFromInt(booking) == models.Rented {
		return false
	}
	database.BookRoom(roomNumber)
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

func CreateRoomWithMessage(room models.Room) string {
	success := database.CreateRoom(room)
	if !success {
		return fmt.Sprintf("Falha ao criar quarto com o número: %d.", room.Number)
	}
	return fmt.Sprintf("Quarto de número: %d adicionado com sucesso!", room.Number)
}

func DeleteRoomWithMessage(id int) string {
	success := database.DeleteRoom(id)
	if !success {
		return fmt.Sprintf("Falha ao apagar o quarto com o id: %d", id)
	}
	return fmt.Sprintf("Sucesso ao apagar o quarto com o id: %d", id)
}

func BookRoomWithMessage(number int) string {
	success := BookRoom(number)
	if !success {
		return fmt.Sprintf("Falha ao reservar o quarto com o número, esse quarto já está reservado: %d", number)
	}
	return fmt.Sprintf("Sucesso ao reservar o quarto com o número: %d", number)
}

func UnBookRoomWithMessage(number int) string {
	success := UnbookRoom(number)
	if !success {
		return fmt.Sprintf("Falha ao liberar o quarto com o número, esse quarto já está reservado: %d", number)
	}
	return fmt.Sprintf("Sucesso ao liberar o quarto com o número: %d", number)
}
