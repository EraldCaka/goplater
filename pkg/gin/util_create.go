package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateUtilEnvs(directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "util.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating util.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(`package util

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
`)
	if err != nil {
		fmt.Println("Error writing to util.go file:", err)
		return err
	}
	fmt.Println("util.go file created successfully")
	return nil
}

func CreatePasswordHasher(directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "bcrypt.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating bcrypt.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(`package util

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
`)
	if err != nil {
		fmt.Println("Error writing to bcrypt.go file:", err)
		return err
	}
	fmt.Println("bcrypt.go file created successfully")
	return nil
}
