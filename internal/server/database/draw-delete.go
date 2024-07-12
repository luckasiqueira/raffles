package database

import (
	"net/http"
	"raffles/utils/logger"
)

// DrawDelete connects to DB and remove the draw for the given number
func DrawDelete(number int) int {
	db := Connect()
	_, err := db.Exec("DELETE FROM `participants` WHERE `Draw` = ?;", number)
	if err != nil {
		logger.NewLog(err.Error())
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusNoContent
}
