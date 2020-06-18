package logger

import (
	"io"
	l "log"
	"os"
)

// Info holds the logger
var Info *l.Logger

var file *os.File

// Create inits the Logger
func Create() {
	var err error
	file, err = os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic("Cannot open logfile")
	}

	writer := io.MultiWriter(file, os.Stdout)
	Info = l.New(writer, "", l.Ldate|l.Ltime)
}

// Destroy closes writers
func Destroy() {
	file.Close()
}
