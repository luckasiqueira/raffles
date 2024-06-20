package database

import (
	"fmt"
	"net/http"
)

// DrawEdit connects to DB and performs the UPDATE of info for the given older draw number
func DrawEdit(oldNumber, number int, name, contact string) int {
	db := Connect()
	_, err := db.Exec("UPDATE `participants` SET `Name` = ?, `Draw` = ?, `Contact` = ? WHERE `Draw` = ?;", name, number, contact, oldNumber)
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusAccepted
}
