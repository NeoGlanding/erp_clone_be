package customer_repository

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
)

func GetCustomerByIdAndPartyId(id string, party_id string) (models.Customer, bool) {
	var data models.Customer

	result := db.PSQL.Table("customers").Where("id = ? AND party_id = ?", id, party_id).
		Preload("Addresses").
		Preload("Contacts").
		Preload("Party").
		Preload("CustomerType").
		Preload("CustomerPartnership").
		Preload("Party.File").
		Preload("Party.Country").
		Preload("File").
		Preload("Country").
		Find(&data)

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

func GetCustomerContactById(customerContactId string, customerId string) (models.CustomerContact, bool) {

	var data models.CustomerContact

	result := db.PSQL.Table("customer_contacts").Where("id = ? AND customer_id = ?", customerContactId, customerId).Find(&data)

	if result.RowsAffected == 0 {
		return data, false
	}

	return data, true

}
