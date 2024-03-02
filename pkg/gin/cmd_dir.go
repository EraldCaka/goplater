package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateMainGo(username, projectName, directory string) error {
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
	_, err = file.WriteString(fmt.Sprintf(`package main

import (
	"github.com/%s/%s/router"
	"github.com/%s/%s/util"
)

func main() {
	util.InitEnvironmentVariables()
	router.InitRouter()
	router.Start("0.0.0.0:5000")
}
`, username, projectName, username, projectName))
	if err != nil {
		fmt.Println("Error writing to main.go file:", err)
		return err
	}
	fmt.Println("main.go file created successfully")
	return nil
}
