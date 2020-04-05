package utils

import (
	"io"
	"log"
	"os"
)

// LogFatal log the specified error to a log file and quit the program
func LogFatal(err error) {
	logFile, err1 := os.Create("errorLog.txt")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer logFile.Close()

	_, err2 := io.WriteString(logFile, err.Error())
	if err2 != nil {
		log.Fatal(err2)
	}

	log.Fatal(err)
}
