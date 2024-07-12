package logger

import (
	"fmt"
	"log"
	"os"
)

func NewLog(message string) {
	file, err := os.OpenFile("errors.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Print("Error while loading log file")
	}
	defer file.Close()
	log.SetOutput(file)
	log.Println(message)
}
