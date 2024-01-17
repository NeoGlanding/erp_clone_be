package main

import (
	"fmt"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/initializers"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if (err != nil) {
		fmt.Println("Error loading .env",err)
	}

	initializers.PSQLInit()
}

func main() {
	db.PSQLMigrate()
}