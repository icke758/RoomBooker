package services

import (
	"fmt"
	"room/database"
	"room/models"
)

func BookRoom(roomNumber int) {
	booking := database.CheckBooking(roomNumber)
	if models.AvaliabilityStatusFromInt(booking) == models.Available {
		database.BookRoom(roomNumber)
	} else {
		fmt.Println("Esse quarto já está reservado")
	}
}

func UnbookRoom(roomNumber int) {
	booking := database.CheckBooking(roomNumber)
	if models.AvaliabilityStatusFromInt(booking) == models.Rented {
		database.UnbookRoom(roomNumber)
	} else {
		fmt.Println("Esse quarto já está liberado")
	}
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
