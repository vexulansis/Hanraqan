package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Подключение .env и изъятие значения по ключу
func GetEnvStr(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}
