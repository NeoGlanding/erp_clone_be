package userpartypermissions

import (
	"strings"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func RetrievePermission(partyId string, userId string) (string, bool) {
	var data models.UserPartyPermission;
	query := db.PSQL.Table("user_party_permissions").Select("permission").Where("party_id = ? AND user_id = ?", partyId, userId).Find(&data)


	if query.RowsAffected == 0 {
		return "", false
	} else {
		return strings.ToLower(string(data.Permission)), true
	}
}