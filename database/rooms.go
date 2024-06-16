package database

import (
	"fmt"
	"log"
	"room/models"
)

func ShowRooms() string {
	db := openDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM rooms;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var result string
	for rows.Next() {
		var room models.Room
		var avaliability int

		err = rows.Scan(&room.ID, &room.Number, &avaliability, &room.Price)
		if err != nil {
			log.Fatal(err)
		}
		room.Avaliability = models.AvaliabilityStatusFromInt(avaliability)
		result += fmt.Sprintf("Quarto: %d, Disponibilidade: %s, Preço: %.2f\n", room.Number, room.Avaliability.ToString(), room.Price)
	}
	return result
}

func CreateRoom(room models.Room) {
	db := openDatabase()
	defer db.Close()

	availability := room.Avaliability.ToInt()

	_, err := db.Exec("INSERT INTO rooms (number, avaliability, daily) VALUES (?, ?, ?);", room.Number, availability, room.Price)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", room.Number, "criado")
}

func DeleteRoom(id int) {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("DELETE FROM rooms WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto com o Id:", id, "deletado")
}
