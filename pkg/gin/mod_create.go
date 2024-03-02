package gin

import (
	"fmt"
	"os/exec"
)

func CreateGoMod() error {
	cmdGetGin := exec.Command("go", "get", "github.com/gin-gonic/gin")
	if err := cmdGetGin.Run(); err != nil {
		fmt.Println("Error installing Gin:", err)
		return err
	}
	fmt.Println("25% completed")
	cmdGetCors := exec.Command("go", "get", "github.com/gin-contrib/cors")
	if err := cmdGetCors.Run(); err != nil {
		fmt.Println("Error installing Gin Cors:", err)
		return err
	}
	fmt.Println("50% completed")
	cmdGetPgx := exec.Command("go", "get", "github.com/jackc/pgx/v5")
	if err := cmdGetPgx.Run(); err != nil {
		fmt.Println("Error installing pgx:", err)
		return err
	}
	fmt.Println("75% completed")
	cmdGetEnv := exec.Command("go", "get", "github.com/joho/godotenv")
	if err := cmdGetEnv.Run(); err != nil {
		fmt.Println("Error installing godotenv:", err)
		return err
	}
	fmt.Println("Go module created successfully")
	return nil
}
