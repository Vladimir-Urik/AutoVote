package logger

import (
	"fmt"
	"log"
	"os"
)

func SetupLogging() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

// Info logs a message at the info level
func Info(message string) {
	log.Println("[INFO] " + message)
	fmt.Println("[INFO] " + message)
}

// InfoFm logs a message at the info level with a format
func InfoFm(message string, args ...interface{}) {
	log.Printf("[INFO] "+message, args...)
	fmt.Printf("[INFO] "+message, args...)
}

// Debug logs a message at the debug level
func Debug(message string) {
	log.Println("[DEBUG] " + message)
	fmt.Println("[DEBUG] " + message)
}

// Error logs a message at the error level
func Error(message string) {
	log.Println("[ERROR] " + message)
	fmt.Println("[ERROR] " + message)
}

// Fatal logs a message at the fatal level
func Fatal(message string) {
	log.Println("[FATAL] " + message)
	fmt.Println("[FATAL] " + message)
}

// Warn logs a message at the warn level
func Warn(message string) {
	log.Println("[WARN] " + message)
	fmt.Println("[WARN] " + message)
}
