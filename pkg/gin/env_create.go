package gin

import (
	"fmt"
	"os"
)

func CreateDotEnv(projectName string) error {
	file, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("DB_URL=postgres://postgres:1234@db:5432/%v\nDB_URL_MIGRATIONS=postgres://postgres:1234@localhost:5435/%v?host=/var/run/postgresql/data&sslmode=disable", projectName, projectName))
	if err != nil {
		fmt.Println("Error writing to .env file:", err)
		return err
	}
	fmt.Println(".env file created successfully")
	return nil
}
