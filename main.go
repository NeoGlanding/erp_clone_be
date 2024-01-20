package main

import (
	"fmt"
	"net/http"

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
}

func main() {
	db.PSQLMigrate()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "storm"})
	})

	routes.Auth(r)

	r.Run()

}