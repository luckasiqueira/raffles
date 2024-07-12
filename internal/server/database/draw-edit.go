package database

import (
	"net/http"
	"raffles/utils/logger"
)

// DrawEdit connects to DB and performs the UPDATE of info for the given older draw number
func DrawEdit(oldNumber, number int, name, contact string) int {
	db := Connect()
	_, err := db.Exec("UPDATE `participants` SET `Name` = ?, `Draw` = ?, `Contact` = ? WHERE `Draw` = ?;", name, number, contact, oldNumber)
	if err != nil {
		logger.NewLog(err.Error())
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusAccepted
}
