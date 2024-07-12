package info

import (
	"github.com/joho/godotenv"
	"raffles/utils/logger"
)

var Env = info()

// info reads our .env file, and get every key|value as map of strings
func info() map[string]string {
	data, err := godotenv.Read()
	if err != nil {
		logger.NewLog(err.Error())
	}
	return data
}
