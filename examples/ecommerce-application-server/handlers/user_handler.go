package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/EraldCaka/ecommerce-application-server/db"
	"github.com/EraldCaka/ecommerce-application-server/types"
)

func GetAllUsers(ctx *gin.Context, dbConn *db.Postgres) {
	users, err := dbConn.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to retrieve users from DB"})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func GetUserByID(ctx *gin.Context, dbConn *db.Postgres) {
	email := ctx.Param("email")
	user := dbConn.GetUserByID(ctx, email)
	if user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context, dbConn *db.Postgres) {
	var user types.UserCreate
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid request payload"})
		return
	}
	userID, err := dbConn.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "fail", "message": "Failed to create user"})
		return
	}
	ctx.JSON(http.StatusCreated, userID)
}
