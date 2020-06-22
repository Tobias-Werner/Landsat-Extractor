package logger

import (
	"io"
	"log"
	"os"
	"sync"
)

var logger *log.Logger = create()
var file *os.File
var lock sync.Mutex

// Create inits the Logger
func create() *log.Logger {
	var err error
	file, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Cannot open logfile")
	}

	writer := io.MultiWriter(file, os.Stdout)
	return log.New(writer, "", log.Ltime)
}

// Info prints debug messages
func Info(msg string) {
	lock.Lock()
	logger.Println(msg)
	lock.Unlock()
}

// Destroy closes writers
func Destroy() {
	file.Close()
}
