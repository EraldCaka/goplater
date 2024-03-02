package gin

import (
	"fmt"
	"os"
)

func CreateGitignore() error {
	file, err := os.Create(".gitignore")
	if err != nil {
		fmt.Println("Error creating .gitignore file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString("*.env\n/bin\n.idea")
	if err != nil {
		fmt.Println("Error writing to .gitignore file:", err)
		return err
	}
	fmt.Println(".gitignore file created successfully")
	return nil
}
