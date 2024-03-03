package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateRouter(username, projectName, directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "router.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating router.go file:", err)
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf(`package router

import (
	"context"
	"github.com/%s/%s/db"
	"github.com/%s/%s/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

`, username, projectName, username, projectName) + "var r *gin.Engine\n\nfunc InitRouter() {\n\tr = gin.Default()\n\tdbConn, err := db.NewPGInstance(context.Background())\n\tif err != nil {\n\t\tlog.Fatalf(\"could not initialize database connection: %s\", err)\n\t\treturn\n\t}\n\n\tr.Use(cors.New(cors.Config{\n\t\tAllowOrigins:     []string{\"http://localhost:5173\"},\n\t\tAllowMethods:     []string{\"GET\", \"POST\"},\n\t\tAllowHeaders:     []string{\"Content-Type\"},\n\t\tExposeHeaders:    []string{\"Content-Length\"},\n\t\tAllowCredentials: true,\n\t\tAllowOriginFunc: func(origin string) bool {\n\t\t\treturn true\n\t\t},\n\t\tMaxAge: 12 * time.Hour,\n\t}))\n\n\tr.GET(\"/users/:userID\", func(ctx *gin.Context) {\n\t\thandlers.GetUserByID(ctx, dbConn)\n\t})\n\tr.GET(\"/users\", func(ctx *gin.Context) {\n\t\thandlers.GetAllUsers(ctx, dbConn)\n\t})\n\tr.POST(\"/users\", func(ctx *gin.Context) {\n\t\thandlers.CreateUser(ctx, dbConn)\n\t})\n}\n\nfunc Start(addr string) error {\n\treturn r.Run(addr)\n}")
	if err != nil {
		fmt.Println("Error writing to router.go file:", err)
		return err
	}
	fmt.Println("router.go file created successfully")
	return nil
}
