package db

import (
	"github.com/automa8e_clone/models"
	"gorm.io/gorm"
)

var PSQL *gorm.DB;

func PSQLMigrate() {
	PSQL.AutoMigrate(&models.User{})
}