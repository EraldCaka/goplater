package empty

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateMainGo(directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "main.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating main.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString("package main\n\nimport \"fmt\"\n\nfunc main() {\n\tfmt.Println(\"Generated with goplater\")\n}\n")
	if err != nil {
		fmt.Println("Error writing to main.go file:", err)
		return err
	}
	fmt.Println("main.go file created successfully")
	return nil
}
