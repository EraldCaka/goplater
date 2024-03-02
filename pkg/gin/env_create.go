package gin

import (
	"fmt"
	"os"
)

func CreateDotEnv() error {
	file, err := os.Create(".env")
	if err != nil {
		fmt.Println("Error creating .env file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString("DB_URL=\nSECRET=")
	if err != nil {
		fmt.Println("Error writing to .env file:", err)
		return err
	}
	fmt.Println(".env file created successfully")
	return nil
}
