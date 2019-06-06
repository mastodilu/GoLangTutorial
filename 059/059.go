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

func sayHi() {
	fmt.Println("Hi!")
}

func tryLogFatal() {
	defer sayHi()
	_, err := os.Open("delete-me.txt")
	if err != nil {
		// stampa nel log e termina la goroutine
		// 💥 senza nemmeno chiamare le defer
		log.Fatal("File can't be opened.")
	}
}

func tryLogPanic() {
	defer sayHi()
	_, err := os.Open("delete-me.txt")
	if err != nil {
		// stampa nel log e termina la goroutine
		log.Panic("File can't be opened.")
	}
}

func main() {
	fmt.Println("Log and error handling.")

	tryLog()

	sendOutputToLogFile() // 💥

	tryLogPanic() // 💥

	tryLogFatal() // 💥

	fmt.Println("BYE")
}
