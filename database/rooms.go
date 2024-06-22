package database

import (
	"database/sql"
	"log"
	"room/models"
)

func FindRooms() (*sql.Rows, error) {
	db := openDatabase()
	rows, err := db.Query("SELECT number, avaliability, daily, period FROM rooms;")
	if err != nil {
		db.Close()
		return nil, err
	}
	return rows, nil
}

func FindRoomByNumber(roomNumber int) (models.Room, error) {
	db := openDatabase()

	var room models.Room
	var availability int

	err := db.QueryRow("SELECT number, avaliability, daily, period FROM rooms WHERE number = ?;", roomNumber).Scan(&room.Number, &availability, &room.Price, &room.Period)
	if err != nil {
		db.Close()
		return room, err
	}

	room.Avaliability = models.AvaliabilityStatusFromInt(availability)

	return room, nil
}

func CreateRoom(room models.Room) bool {
	db := openDatabase()
	defer db.Close()

	availability := room.Avaliability.ToInt()

	_, err := db.Exec("INSERT INTO rooms (number, avaliability, daily, period) VALUES (?, ?, ?, ?);", room.Number, availability, room.Price, room.Period)
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
