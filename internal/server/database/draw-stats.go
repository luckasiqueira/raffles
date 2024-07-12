package database

import (
	"raffles/utils/info"
	"raffles/utils/logger"
	"strconv"
)

// DrawStatus connects to DB to check the raffle status and stats of participants
func DrawStatus() Status {
	db := Connect()
	var status Status
	var done int
	err := db.QueryRow("SELECT COUNT(DISTINCT `Name`) AS unique_name_count FROM `participants`;").Scan(&status.Participants)
	if err != nil {
		logger.NewLog(err.Error())
	}
	err = db.QueryRow("SELECT COUNT(*) FROM `participants`;").Scan(&status.Bought)
	if err != nil {
		logger.NewLog(err.Error())
	}
	n, err := strconv.Atoi(info.Env["TOTAL_NUMS"])
	if err != nil {
		logger.NewLog(err.Error())
	}
	status.Available = 50 - n
	err = db.QueryRow("SELECT `Status` FROM `status` WHERE 1").Scan(&done)
	if err != nil {
		logger.NewLog(err.Error())
	}
	if done == 1 {
		status.Done = true
		err = db.QueryRow("SELECT `Draw`, `Name` FROM `status` WHERE 1").Scan(&status.WinnerDraw, &status.WinnerName)
		if err != nil {
			logger.NewLog(err.Error())
		}
	}
	defer db.Close()
	return status
}
