package helpers

import (
	"log"
	"os"
)
// Инфо-логгер
func InfoLog(message string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger := log.New(file, "INFO: ", log.Ldate|log.Ltime)
	InfoLogger.Println(message)
} 
// Логгер ошибок
func ErrorLog(message string) {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ErrorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime)
	ErrorLogger.Println(message)
}
