package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/automa8e_clone/config"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if (err != nil) {
		fmt.Println("Error loading .env",err)
	}

	initializers.PSQLInit()
	db.PSQLSeed()

	initializers.Validator()

	config.AppConfig.JWT_SECRET = os.Getenv("JWT_SECRET")
	config.AppConfig.JWT_TOKEN_VERSION = os.Getenv("JWT_TOKEN_VERSION")
	config.AppConfig.JWT_TIME_EXPIRATION = os.Getenv("JWT_TOKEN_EXPIRATION_MINUTE")

}

func main() {
	db.PSQLMigrate()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "storm"})
	})

	routes.Auth(r)
	routes.Party(r)
	routes.Country(r)
	routes.Users(r)

	r.Run()

}