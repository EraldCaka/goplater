package gin

import (
	"fmt"
	"github.com/EraldCaka/goplater/pkg/dir"
	"os"
	"path/filepath"
)

func CreateUserHandler(username, projectName, directory string) error {
	if err := dir.CreateDir(directory); err != nil {
		return err
	}
	filePath := filepath.Join(directory, "user_handler.go")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating user_handler.go file:", err)
		return err
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf(`package main

import (
	"github.com/%s/%s/router"
	"github.com/%s/%s/util"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/%s/%s/db"
)

func GetAllUsers(ctx *gin.Context, dbConn *db.Postgres) {
	users := dbConn.GetAllUsers(ctx)
	if users == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to retrieve users from DB"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": users})
}

func GetUserByEmail(ctx *gin.Context, dbConn *db.Postgres) {
	email := ctx.Param("email")
	user := dbConn.GetUserByEmail(ctx, email)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}

func CreateUser(ctx *gin.Context, dbConn *db.Postgres) {
	var user UserCreate
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid request payload"})
		return
	}
	err := dbConn.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "user": user})
}


func DeleteUser(ctx *gin.Context, dbConn *db.Postgres) {
	userID := ctx.Param("userID")
	err := dbConn.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to delete user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "User deleted successfully"})
}

func UpdateUser(ctx *gin.Context, dbConn *db.Postgres) {
	userID := ctx.Param("userID")
	var user UserUpdate
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid request payload"})
		return
	}
	err := dbConn.UpdateUser(ctx, userID, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to update user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": user})
}
`, username, projectName, username, projectName, username, projectName))

	if err != nil {
		fmt.Println("Error writing to user_handler.go file:", err)
		return err
	}
	fmt.Println("user_handler.go file created successfully")
	return nil
}
