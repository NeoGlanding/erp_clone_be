package users_repository

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func FindByEmail(email string, bind *models.User) bool {
	db := db.PSQL.Table("users").Where("email = ? AND deleted_at IS NULL", email).Find(&bind)
	return db.RowsAffected != 0;
}