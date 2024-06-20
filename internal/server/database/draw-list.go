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

// DrawSingle receives a number, and collects participant's data of this number
func DrawSingle(number int) Participant {
	var p Participant
	db := Connect()
	err := db.QueryRow("SELECT * FROM `participants` WHERE `Draw` = ?;", number).Scan(&p.Name, &p.Draw, &p.Contact)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	return p
}
