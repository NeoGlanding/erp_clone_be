package files

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func GetFileById(id string, user_id string) (models.File, bool) {
	var data models.File



	result := db.PSQL.Table("files").Where("id = ? AND user_id = ?", id, user_id).Find(&data)

	if result.RowsAffected == 0 {
		return data, false
	}

	return data, true
}