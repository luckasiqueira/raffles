package database

import (
	"net/http"
	"raffles/utils/logger"
)

// DrawAdd connects do DB to create new participant's draws
func DrawAdd(number int, name, contact string) int {
	db := Connect()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM `participants` WHERE `Draw` = ?;", number).Scan(&count)
	if err != nil {
		logger.NewLog(err.Error())
		return http.StatusInternalServerError
	}
	if count > 0 {
		return http.StatusConflict
	}
	_, err = db.Exec("INSERT INTO `participants` (`Name`, `Draw`, `Contact`) VALUES (?, ?, ?);", name, number, contact)
	if err != nil {
		logger.NewLog(err.Error())
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusCreated
}
