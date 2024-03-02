package main

import (
	"bufio"
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"github.com/EraldCaka/goplater/pkg/empty"
	"github.com/EraldCaka/goplater/pkg/gin"
	"os"
)

func main() {
	fmt.Println("\n _____ _____ _____ __    _____ _____ _____ _____ \n|   __|     |  _  |  |  |  _  |_   _|   __| __  |\n|  |  |  |  |   __|  |__|     | | | |   __|    -|\n|_____|_____|__|  |_____|__|__| |_| |_____|__|__|\n\n")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Your github name: ")
	githubUser := dir.CliScan(scanner)

	fmt.Print("Your project name : ")
	projectName := dir.CliScan(scanner)

	if err := dir.CreateDir(projectName); err != nil {
		return
	}
	dirUrl := fmt.Sprintf("./%v", projectName)
	if err := dir.CreateGoMod(githubUser, projectName, dirUrl); err != nil {
		return
	}
	fmt.Println("\n-------------------------------")
	fmt.Println("1 - empty project\n2 - gin project\n3 - fiber project\n4 - echo project\n5 - mux project\n6 - standard http")
	fmt.Println("Select the project template you want to create:")
	choice := dir.CliScan(scanner)
	switch choice {
	case "1":
		if err := empty.CreateMainGo("cmd"); err != nil {
			return
		}
		break
	case "2":
		if err := gin.CreateGoMod(); err != nil {
			return
		}
		if err := gin.CreateGitignore(); err != nil {
			return
		}
		if err := gin.CreateDotEnv(); err != nil {
			return
		}
		break
	}

}
