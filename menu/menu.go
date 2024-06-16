package menu

import (
	"fmt"
	"room/models"
	"room/utils"
)

func ShowMenuOptions() {
	menuMap := map[string]string{
		"1": "Mostrar quartos",
		"2": "Reservar quarto",
		"3": "Disponibilizar quarto",
		"4": "Criar quarto",
		"5": "Deletar quarto",
	}

	keys := utils.SortKeys(menuMap)

	for _, key := range keys {
		fmt.Printf("[%s] -> [%s] \n", key, menuMap[key])
	}
	fmt.Println("-----------------------------")
}

func ShowCreateRoomMenu() models.Room {
	var roomNumber int
	var roomAvailability int
	var price float64

	fmt.Println("Criar quarto")
	fmt.Println("Numero do quarto: ")
	fmt.Scanf("%d", &roomNumber)
	fmt.Println("Disponibilidade: ")
	fmt.Scanf("%d", &roomAvailability)
	fmt.Println("Preço diário: ")
	fmt.Scanf("%f", &price)

	availability := models.AvaliabilityStatusFromInt(roomAvailability)

	return models.Room{
		Number:       roomNumber,
		Avaliability: availability,
		Price:        float32(price),
	}
}

func ShowRoomsMenu(rooms string) {
	fmt.Println("Mostrar quartos")
	fmt.Println(rooms)
}

func ShowRoomsToBook() {
	fmt.Println("Reservar quarto")
	fmt.Println("Quarto a ser reservado?")
}

func ShowRoomsToUnBook() {
	fmt.Println("Reservar quarto")
	fmt.Println("Quarto a ser liberado: ")
}

func ShowDeleteText() {
	fmt.Println("Apagar Quarto")
	fmt.Println("Quarto a ser apagado: ")
}
