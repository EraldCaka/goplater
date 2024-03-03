package router

import (
	"context"
	"github.com/EraldCaka/gin-project-example/db"
	"github.com/EraldCaka/gin-project-example/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

var r *gin.Engine

func InitRouter() {
	r = gin.Default()
	dbConn, err := db.NewPGInstance(context.Background())
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
		return
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/users/:userID", func(ctx *gin.Context) {
		handlers.GetUserByID(ctx, dbConn)
	})
	r.GET("/users", func(ctx *gin.Context) {
		handlers.GetAllUsers(ctx, dbConn)
	})
	r.POST("/users", func(ctx *gin.Context) {
		handlers.CreateUser(ctx, dbConn)
	})
}

func Start(addr string) error {
	return r.Run(addr)
}