package main

import (
	"fmt"
	"log"
	"os"
)

func tryLog() {
	//stampa nel terminale un messaggio di log
	_, err := os.Open("delete-me.txt")
	if err != nil {
		// fmt.Println("File can't be opened.")
		log.Println("File can't be opened.", err)
	}
}

func sendOutputToLogFile() {
	logFile, err := os.Create("log.txt")
	if err != nil {
		fmt.Println("File can't be opened.")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	file, err := os.Open("delete-me.txt")
	if err != nil {
		log.Println("File can't be opened.", err)
	}
	defer file.Close()

	fmt.Println("Check the log file.")
}

func tryLogFatal() {
	_, err := os.Open("delete-me.txt")
	if err != nil {
		// stampa nel log e termina
		// 💥 senza nemmeno chiamare le defer
		log.Fatal("File can't be opened.")
	}
}

func main() {
	fmt.Println("Log and error handling.")

	tryLog()

	sendOutputToLogFile() // 💥

	tryLogFatal() // 💥
}
