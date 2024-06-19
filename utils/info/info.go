package info

import (
	"fmt"
	"github.com/joho/godotenv"
)

var Env = info()

// info reads our .env file, and get every key|value as map of strings
func info() map[string]string {
	data, err := godotenv.Read()
	if err != nil {
		fmt.Println(err.Error())
	}
	return data
}
