package database

import (
	"fmt"
)

// DrawList connects to DB and get details of all participants
func DrawList() []Participant {
	db := Connect()
	var draws []Participant
	rows, err := db.Query("SELECT * FROM `participants`;")
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var draw Participant
		err = rows.Scan(&draw.Name, &draw.Draw, &draw.Contact)
		if err != nil {
			fmt.Println(err.Error())
		}
		draws = append(draws, draw)
	}
	defer db.Close()
	return draws
}
