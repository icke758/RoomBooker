package menu

import (
	"fmt"
	"log"
	"room/database"
	"room/models"
	"room/services"
	"room/utils"
)

func ShowMenuOptions() {
	menuMap := map[string]string{
		"1": "Mostrar quartos",
		"2": "Reservar quarto",
		"3": "Disponibilizar quarto",
		"4": "Criar quarto",
		"5": "Apagar quarto",
	}

	keys := utils.SortKeys(menuMap)

	for _, key := range keys {
		fmt.Printf("[%s] -> [%s] \n", key, menuMap[key])
	}
	fmt.Println("-----------------------------")
}

func ShowRooms() string {
	rows, err := database.FindRooms()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var room models.Room
		var availability int
		err = rows.Scan(&room.Number, &availability, &room.Price)
		if err != nil {
			log.Fatal(err)
		}
		room.Avaliability = models.AvaliabilityStatusFromInt(availability)
		result += fmt.Sprintf("Quarto: %d, Disponibilidade: %s, Preço: %.2f\n", room.Number, room.Avaliability.ToString(), room.Price)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func ShowCreateRoomMenu() models.Room {
	var roomNumber int
	var roomAvailability int
	var daily float64

	fmt.Println("Criar quarto")
	fmt.Println("Numero do quarto: ")
	fmt.Scanf("%d", &roomNumber)
	fmt.Printf("Disponibilidade (%d para alugado, %d para disponivel): ", models.Rented.ToInt(), models.Available.ToInt())
	fmt.Scanf("%d", &roomAvailability)
	fmt.Println("Preço diário: ")
	fmt.Scanf("%f", &daily)

	availability := models.AvaliabilityStatusFromInt(roomAvailability)

	room := models.Room{
		Number:       roomNumber,
		Avaliability: availability,
		Price:        float32(daily),
	}

	return room
}

func ShowRoomsMenu(rooms string) {
	fmt.Println("Mostrar quartos")
	fmt.Println(rooms)
}

func ShowRoomsToBook() {
	fmt.Println("Reservar quarto")
	fmt.Println("Quarto a ser reservado: ")
}

func ShowRoomsToUnBook() {
	fmt.Println("Liberar quarto")
	fmt.Println("Quarto a ser liberado: ")
}

func ShowDeleteText() {
	fmt.Println("Apagar Quarto")
	fmt.Println("Quarto a ser apagado: ")
}

func CreatedRoomMessage(room models.Room) {
	message := services.CreateRoomWithMessage(room)
	fmt.Println(message)
}

func DeletedRoomMessage(id int) {
	message := services.DeleteRoomWithMessage(id)
	fmt.Println(message)
}

func BookRoomMessage(number int) {
	message := services.BookRoomWithMessage(number)
	fmt.Println(message)
}

func UnBookRoomMessage(number int) {
	message := services.UnBookRoomWithMessage(number)
	fmt.Println(message)
}
