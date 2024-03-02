package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateUserType(directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "User.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating User.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString("package types\n\ntype User struct {\n\tID       int    `json:\"_id,omitempty\"`\n\tUsername string `json:\"username\"`\n\tPassword string `json:\"password,omitempty\"`\n\tEmail    string `json:\"email\"`\n}\ntype UserUpdate struct {\n\tUsername string `json:\"username\"`\n\tPassword string `json:\"password,omitempty\"`\n\tEmail    string `json:\"email\"`\n}\ntype UserCreate struct {\n\tUsername string `json:\"username\"`\n\tPassword string `json:\"password,omitempty\"`\n\tEmail    string `json:\"email\"`\n}")
	if err != nil {
		fmt.Println("Error writing to User.go file:", err)
		return err
	}
	fmt.Println("User.go file created successfully")
	return nil
}
