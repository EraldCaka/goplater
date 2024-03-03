package util

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	DB_URL string
	SECRET string
)

func InitEnvironmentVariables() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	DB_URL = os.Getenv("DB_URL")
	SECRET = os.Getenv("SECRET")
	func() {
		executeMigrateScript := exec.Command("make", "migrate")
		if err := executeMigrateScript.Run(); err != nil {
			fmt.Println("Error executing migration scripts:", err)
			return
		}
	}()
}
