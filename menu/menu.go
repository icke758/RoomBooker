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
		"6": "Mostrar recibo",
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
		err = rows.Scan(&room.Number, &availability, &room.Price, &room.Period)
		if err != nil {
			log.Fatal(err)
		}
		room.Avaliability = models.AvaliabilityStatusFromInt(availability)
		result += fmt.Sprintf("Quarto: %d, Disponibilidade: %s, Preço: %.2f, Período da reserva: %d\n", room.Number, room.Avaliability.ToString(), room.Price, room.Period)
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

func ShowReceiptMenu() int {
	var roomNumber int

	fmt.Println("Digite o número do quaro que você deseja ver o recibo: ")
	fmt.Scanf("%d", &roomNumber)

	return roomNumber
}

func ShowReceipt(roomNumber int) {
	fmt.Println("Mostrar recibo")
	receiptData := services.CreateReceipt(roomNumber)
	fmt.Println("		_._._._ ROOM BOOKER _._._._")
	fmt.Println("		-------    Recibo   -------")
	fmt.Printf("Número: 			%d\n", receiptData.Number)
	fmt.Printf("Disponibilidade: 		%s\n", receiptData.Avaliability.ToString())
	fmt.Printf("Período: 			%d dias\n", receiptData.Period)
	fmt.Printf("Preço diário: 			%.2f\n", receiptData.Price)
	fmt.Printf("Preço total: 			%.2f\n", receiptData.TotalAmount)
	fmt.Println("----------------------------------")
}

func ShowRoomsMenu(rooms string) {
	fmt.Println("Mostrar quartos")
	fmt.Println(rooms)
}

func ShowRoomsToBook() int {
	var roomNumber int
	fmt.Println("Reservar quarto")
	fmt.Println("Quarto a ser reservado: ")
	fmt.Scanf("%d", &roomNumber)

	return roomNumber
}

func ShowBookingPeriod() int {
	var period int
	fmt.Println("Período da reserva: ")
	fmt.Scanf("%d", &period)

	return period
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
	message := services.CreateRoomLog(room)
	fmt.Println(message)
}

func DeletedRoomMessage(id int) {
	message := services.DeleteRoomLog(id)
	fmt.Println(message)
}

func BookRoomMessage(number int, period int) {
	message := services.BookRoomLog(number, period)
	fmt.Println(message)
}

func UnBookRoomMessage(number int) {
	message := services.UnBookRoomWithMessage(number)
	fmt.Println(message)
}
