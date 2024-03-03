package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DB_URL string
)

func InitEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
}
