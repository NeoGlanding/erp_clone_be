package initializers

import (
	"fmt"
	"os"

	"github.com/automa8e_clone/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func PSQLInit() {
	dsn := os.Getenv("PSQL_GORM")
	psql, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	fmt.Println(err)
	if (err != nil) {
		panic(err)
	}

	fmt.Println("Success connect to PSQL DB")

	db.PSQL = psql
}