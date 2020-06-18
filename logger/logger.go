package logger

// Create inits the Logger
func Create() {
	println("Creating")
}

// Info logs a message
func Info(msg string) {
	println(msg)
}

// Destroy closes writers
func Destroy() {
	println("Closing")
}
