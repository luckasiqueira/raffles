package database

import (
	"fmt"
	"net/http"
)

// DrawDelete connects to DB and remove the draw for the given number
func DrawDelete(number int) int {
	db := Connect()
	_, err := db.Exec("DELETE FROM `buyers` WHERE `Draw` = ?;", number)
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusNoContent
}
