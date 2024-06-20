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
	db.Close()
	return roomAvailabilityInt
}

func BookRoom(number int, period int) bool {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET avaliability = 1, period = ? WHERE number = ?;", number, period)
	if err != nil {
		log.Fatal(err)
		return false
	}
	db.Close()
	return true
}

func UnbookRoom(number int) bool {
	db := openDatabase()
	defer db.Close()

	_, err := db.Exec("UPDATE rooms SET avaliability = 2, period = 0 WHERE number = ?;", number)
	if err != nil {
		log.Fatal(err)
		return false
	}
	db.Close()
	return true
}

func FindDailyValueByRoom(number int) (float64, error) {
	db := openDatabase()
	defer db.Close()

	var dailyValue float64

	err := db.QueryRow("SELECT daily FROM rooms WHERE number = ?", number).Scan(&dailyValue)
	if err != nil {
		return 0, err
	}
	db.Close()
	return dailyValue, nil
}

func FindTimePeriodByRoom(number int) (int, error) {
	db := openDatabase()
	defer db.Close()

	var timePeriod int

	err := db.QueryRow("SELECT period FROM rooms WHERE number = ?", number).Scan(&timePeriod)
	if err != nil {
		return 0, err
	}
	db.Close()
	return timePeriod, nil
}
