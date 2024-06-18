package database

import (
	"log"
)

func CheckBooking(number int) int {
	db := openDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT avaliability FROM rooms WHERE number = ?;", number)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var roomAvailabilityInt int

	for rows.Next() {
		err := rows.Scan(&roomAvailabilityInt)
		if err != nil {
			log.Fatal(err)
		}
	}
	return roomAvailabilityInt
}

func BookRoom(number int) bool {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET avaliability = '1' WHERE number = ?;", number)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func UnbookRoom(number int) bool {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET avaliability = 2 WHERE number = ?;", number)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
