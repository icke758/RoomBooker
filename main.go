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
			roomNumber := menu.ShowRoomsToBook()
			period := menu.ShowBookingPeriod()
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
		case 6:
			roomNumber := menu.ShowReceiptMenu()
			menu.ShowReceipt(roomNumber)
		default:
			fmt.Println("Ok, adeus! :D")
			return
		}
	}
}
