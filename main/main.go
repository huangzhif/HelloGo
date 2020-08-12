package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, err := os.OpenFile("./logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Open log file failed,err:", err)
		return
	}

	log.SetOutput(logFile)
	log.Println("test")
}
