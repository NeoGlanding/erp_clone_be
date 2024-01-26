package countries

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func FindById(id string) (models.Country, bool) {
	var data models.Country;
	query := db.PSQL.Table("countries").Where("id = ?", id).Find(&data)
	exist := query.RowsAffected > 0

	return data, exist
}

func FindBy2AlphaCode(code string) (models.Country, bool) {
	var data models.Country;
	query := db.PSQL.Table("countries").Where("two_alpha_code = ?", code).Find(&data)
	exist := query.RowsAffected > 0

	return data, exist
}

func FindBy3AlphaCode(code string) (models.Country, bool) {
	var data models.Country;
	query := db.PSQL.Table("countries").Where("three_alpha_code = ?", code).Find(&data)
	exist := query.RowsAffected > 0

	return data, exist
}