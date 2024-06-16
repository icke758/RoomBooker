package main

import (
	"fmt"
	"room/database"
	"room/menu"
)

func main() {
	var userAction int
	var roomNumber int

	for {
		menu.ShowMenuOptions()

		fmt.Scanf("%d", &userAction)
		switch userAction {
		case 1:
			menu.ShowRoomsMenu(database.ShowRooms())
		case 2:
			menu.ShowRoomsToBook()
			fmt.Scanf("%d", &roomNumber)
			database.BookRoom(roomNumber)
		case 3:
			database.UnbookRoom(roomNumber)
		case 4:
			room := menu.ShowCreateRoomMenu()
			database.CreateRoom(room)
		case 5:
			var id int
			menu.ShowDeleteText()
			fmt.Scanf("%d", &id)
			database.DeleteRoom(id)
		default:
			fmt.Println("Ok, adeus! :D")
			return
		}
	}
}
