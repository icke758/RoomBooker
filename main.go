package main

import (
	"fmt"
	"room/menu"
)

func main() {
	var userAction int

	for {
		menu.ShowMenuOptions()

		fmt.Scanf("%d", &userAction)
		switch userAction {
		case 1:
			rooms := menu.ShowRooms()
			menu.ShowRoomsMenu(rooms)
		case 2:
			roomNumber, period := menu.ShowRoomsToBook()
			menu.BookRoomMessage(roomNumber, period)
		case 3:
			var roomNumber int
			menu.ShowRoomsToUnBook()
			fmt.Scanf("%d", &roomNumber)
			menu.UnBookRoomMessage(roomNumber)
		case 4:
			room := menu.ShowCreateRoomMenu()
			menu.CreatedRoomMessage(room)
		case 5:
			var id int
			menu.ShowDeleteText()
			fmt.Scanf("%d", &id)
			menu.DeletedRoomMessage(id)
		default:
			fmt.Println("Ok, adeus! :D")
			return
		}
	}
}
