package main

import (
	"net/http"

	"github.com/erizkiatama/rasalibrary-backend/auth"
	_authHandler "github.com/erizkiatama/rasalibrary-backend/auth/handler"
	_authRepo "github.com/erizkiatama/rasalibrary-backend/auth/repository"
	"github.com/erizkiatama/rasalibrary-backend/db"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/server/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Server is OK!")
	})

	db := db.NewPostgreSQLDatabase()

	v1 := router.Group("/api/v1")

	authRepo := _authRepo.NewPostgreSQLRepository(db)
	authSvc := auth.NewAuthService(authRepo)
	_authHandler.NewAuthHandler(v1, authSvc)

	router.Run(":8080")
}
