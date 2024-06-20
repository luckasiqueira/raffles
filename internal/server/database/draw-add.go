package database

import "net/http"

// DrawAdd connects do DB to create new participant's draws
func DrawAdd(number int, name, contact string) int {
	db := Connect()
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM `buyers` WHERE `Draw` = ?;", number).Scan(&count)
	if err != nil {
		return http.StatusInternalServerError
	}
	if count > 0 {
		return http.StatusConflict
	}
	_, err = db.Exec("INSERT INTO `buyers` (`Name`, `Draw`, `Contact`) VALUES (?, ?, ?);", name, number, contact)
	if err != nil {
		return http.StatusInternalServerError
	}
	defer db.Close()
	return http.StatusCreated
}
