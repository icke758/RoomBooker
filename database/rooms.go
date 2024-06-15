package database

import (
	"fmt"
	"log"
	"unicode"
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
		var id int
		var room int
		var availability string
		var price float32

		err = rows.Scan(&id, &room, &availability, &price)
		if err != nil {
			log.Fatal(err)
		}
		result += fmt.Sprintf("Quarto: %d, Disponibilidade: %s, Preço: %.2f\n", room, availability, price)
	}
	return result
}

func Clearstring(s string) string {
	var result []rune
	for _, char := range s {
		if unicode.IsLetter(char) {
			result = append(result, char)
		}
	}
	return string(result)
}

func CreateRoom(number int, availability string, daily float64) {
	db := openDatabase()
	defer db.Close()

	Clearstring(availability)

	_, err := db.Exec("INSERT INTO rooms (number, availability, daily) VALUES (?, ?, ?);", number, availability, daily)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", number, "criado")
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
