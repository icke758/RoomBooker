package database

import (
	"fmt"
	"log"
)

func BookRoom(number int) {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET availability = 'alugado' WHERE number = ?;", number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", number, "reservado!")
}

func UnbookRoom(number int) {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET availability = 'disponivel' WHERE number = ?;", number)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Quarto número", number, "disponibilizado")
}
