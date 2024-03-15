package customer_repository

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func GetCustomerByIdAndPartyId(id string, party_id string) (models.Customer, bool) {
	var data models.Customer

	result := db.PSQL.Table("customers").Where("id = ? AND party_id = ?", id, party_id).Find(&data)

	if result.RowsAffected == 0 {
		return data, false
	}

	return data, true
}

func GetCustomerAddressById(customerAddressId string, customerId string) (models.CustomerAddresses, bool) {

	var data models.CustomerAddresses

	result := db.PSQL.Table("customer_addresses").Where("id = ? AND customer_id = ?", customerAddressId, customerId).Find(&data)

	if result.RowsAffected == 0 {
		return data, false
	}

	return data, true

}
