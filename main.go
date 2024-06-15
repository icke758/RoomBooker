package main

import (
	"fmt"
	"room/database"
	"room/menu"
)

func main() {
	var userAction int
	var roomNumber int
	var roomAvaliability string
	var daily float64

	for {
		menu.ShowMenuOptions()

		fmt.Scanf("%d", &userAction)
		if userAction == 1 {
			menu.ShowRoomsMenu(database.ShowRooms())
		} else if userAction == 2 {
			menu.ShowRoomsToBook()
			fmt.Scanf("%d", &roomNumber)
			database.BookRoom(roomNumber)
		} else if userAction == 3 {
			database.UnbookRoom(roomNumber)
		} else if userAction == 4 {
			fmt.Println("Criar quarto")
			fmt.Println("Numero do quarto: ")
			fmt.Scanf("%d", &roomNumber)
			fmt.Println("Disponibilidade: ")
			fmt.Scanf("%s", &roomAvaliability)
			fmt.Println("Preço diário: ")
			fmt.Scanf("%f", &daily)

			database.CreateRoom(roomNumber, roomAvaliability, daily)
		} else if userAction == 5 {
			var id int
			fmt.Println("Apagar Quarto")
			fmt.Println("Qual quarto você quer apagar?")
			fmt.Scanf("%d", &id)

			database.DeleteRoom(id)
		} else {
			fmt.Println("Ok, adeus! :D")
			break
		}
	}
}
