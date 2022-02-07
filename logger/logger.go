package logger

import (
	"fmt"
	"github.com/Vladimir-Urik/AutoVote/managers/config"
	"github.com/Vladimir-Urik/AutoVote/managers/webhook"
	"log"
	"os"
	"time"
)

var (
	messages []string
)

func SetupLogging() {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)
}

func StartWebhook(config *config.Config) {
	go func() {
		for {
			if len(messages) > 0 {
				var build string
				for _, message := range messages {
					build += message + "\n"
				}

				log.Print("[INFO] Sending ", len(messages), " to discord webhook \n")
				fmt.Print("[INFO] Sending ", len(messages), " to discord webhook \n")

				webhook.SendDataToWebhook(build, []webhook.Embed{}, config.LogsWebhook)
			}
			messages = []string{}
			time.Sleep(1 * time.Second)
		}
	}()
}

// Info logs a message at the info level
func Info(message string) {
	log.Println("[INFO] " + message)
	fmt.Println("[INFO] " + message)
	messages = append(messages, "[INFO] "+message)
}

// InfoFm logs a message at the info level with a format
func InfoFm(message string, args ...interface{}) {
	log.Printf("[INFO] "+message, fmt.Sprint(args...))
	fmt.Printf("[INFO] "+message, fmt.Sprint(args...))
	messages = append(messages, "[INFO] "+message+fmt.Sprint(args...))
}

// Debug logs a message at the debug level
func Debug(message string) {
	log.Println("[DEBUG] " + message)
	fmt.Println("[DEBUG] " + message)
	messages = append(messages, "[DEBUG] "+message)
}

// Error logs a message at the error level
func Error(message string) {
	log.Println("[ERROR] " + message)
	fmt.Println("[ERROR] " + message)
	messages = append(messages, "[ERROR] "+message)
}

// Fatal logs a message at the fatal level
func Fatal(message string) {
	log.Println("[FATAL] " + message)
	fmt.Println("[FATAL] " + message)
	messages = append(messages, "[FATAL] "+message)
}

// Warn logs a message at the warn level
func Warn(message string) {
	log.Println("[WARN] " + message)
	fmt.Println("[WARN] " + message)
	messages = append(messages, "[WARN] "+message)
}
