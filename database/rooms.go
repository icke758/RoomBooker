package database

import (
	"database/sql"
	"log"
	"room/models"
)

func FindRooms() (*sql.Rows, error) {
	db := openDatabase()
	rows, err := db.Query("SELECT number, avaliability, daily FROM rooms;")
	if err != nil {
		db.Close()
		return nil, err
	}
	return rows, nil
}

func CreateRoom(room models.Room) bool {
	db := openDatabase()
	defer db.Close()

	availability := room.Avaliability.ToInt()

	_, err := db.Exec("INSERT INTO rooms (number, avaliability, daily) VALUES (?, ?, ?);", room.Number, availability, room.Price)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DeleteRoom(id int) bool {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("DELETE FROM rooms WHERE id = ?", id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
