package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mfduar8766/GoRestAPI/utils"
)

// Log - Logger interface
type Log struct {
	file *os.File
	fileName string
	dir string
}

// LogInstance - Uses as a global for the Log instance
var LogInstance *Log

// CreateLogger - Used to create a log instance
func CreateLogger(dir string, fileName string) {
	log.Print("INFO: CreateLogger()")
	fileDest := filepath.Join(dir, filepath.Base(fileName))
	var fileInstance *os.File
	if _, err := os.Stat(fileDest); os.IsNotExist(err) {
		fileInstance, err := os.Create(fileDest)
		utils.MustNotError(err)
		log.Print("INFO: CreateLogger() Logger Initialized")
		log.SetOutput(fileInstance)
		LogInstance = &Log{
			file: fileInstance,
			fileName: fileName,
			dir: dir,
		}
	}
	fileInstance, err := os.OpenFile(fileDest, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	utils.MustNotError(err)
	log.SetOutput(fileInstance)
	log.Print("INFO: CreateLogger() Logger Already exists")
	LogInstance = &Log{
		file: fileInstance,
		fileName: fileName,
		dir: dir,
	}
}

// Info - General log
func (l *Log) Info(message string) {
	formattedMessage := fmt.Sprintf("INFO: %s", message)
	log.Print(formattedMessage)
	_, err := l.file.WriteString(formattedMessage)
	utils.MustNotError(err)
}

// Panic - Used to log a panic to the console
func (l *Log) Panic(message string) {
	formattedMessage := fmt.Sprintf("PANIC: %s", message)
	log.Panic(formattedMessage)
	_, err := l.file.WriteString(formattedMessage)
	utils.MustNotError(err)
}
