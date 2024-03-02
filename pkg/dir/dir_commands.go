package dir

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func CliScan(scanner *bufio.Scanner) string {
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	return input
}

func CreateDir(dirName string) error {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
		//fmt.Println("Directory created successfully:", dirName)
	} else {
		fmt.Println("Directory already exists:", dirName)
	}
	return nil
}
func CreateGoMod(username, projectName, directory string) error {
	if err := os.Chdir(directory); err != nil {
		fmt.Println("Error changing directory:", err)
		return err
	}
	modFile := fmt.Sprintf("github.com/%s/%s", username, projectName)
	cmd := exec.Command("go", "mod", "init", modFile)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error creating Go module:", err)
		return err
	}
	//fmt.Println("Go module created successfully")
	return nil
}
