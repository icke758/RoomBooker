package services

import (
	"log"
	"room/database"
)

func EvaluateStay(roomNumber int) float64 {
	period, err := database.FindTimePeriodByRoom(roomNumber)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	dailyValue, err := database.FindDailyValueByRoom(roomNumber)
	if err != nil {
		log.Fatal(err)
		return 0
	}
	total := float64(period) * dailyValue
	return total
}
