package controllers

import (
	"errors"
	"fmt"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	customer_repository "github.com/automa8e_clone/repositories/customers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CreateCustomerBody struct {
	Name                       string `json:"name" validate:"required"`
	PartyId                    string `json:"party_id" validate:"required,uuid"`
	BusinessRegistrationNumber string `json:"business_registration_number" validate:"required"`
	Url                        string `json:"url" validate:"required"`
	Remarks                    string `json:"remarks" validate:"required"`
	CustomerTypeId             string `json:"customer_type_id" validate:"required,uuid"`
	CustomerPartnershipId      string `json:"customer_partnership_id" validate:"required,uuid"`
	CountryId                  string `json:"country_id" validate:"required,uuid"`
	FileId                     string `json:"file_id" validate:"required,uuid"`
}

type UpdateCustomerBody struct {
	Name                       string `json:"name" validate:"required"`
	BusinessRegistrationNumber string `json:"business_registration_number" validate:"required"`
	Url                        string `json:"url" validate:"required"`
	Remarks                    string `json:"remarks" validate:"required"`
	CustomerTypeId             string `json:"customer_type_id" validate:"required,uuid"`
	CustomerPartnershipId      string `json:"customer_partnership_id" validate:"required,uuid"`
	CountryId                  string `json:"country_id" validate:"required,uuid"`
	FileId                     string `json:"file_id" validate:"required,uuid"`
}

type CustomerAddressesBody struct {
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2" validate:"required"`
	AddressLine3 string `json:"address_line3"`
	PostalCode   string `json:"postal_code" validate:"required"`
	CountryId    string `json:"country_id" validate:"required,uuid"`
}

type CustomerContactsBody struct {
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type CustomerAddressesBodyPayload struct {
	Addresses []CustomerAddressesBody `json:"addresses" validate:"required"`
}

type CustomerContactsBodyPayload struct {
	Contacts []CustomerContactsBody `json:"contacts" validate:"required,dive,required"`
}

func GetCustomerType(c *gin.Context) {
	data := []models.CustomerType{}
	db.PSQL.Find(&data)
	c.Set("data", data)
	c.Next()
}

func GetCustomerPartnership(c *gin.Context) {
	data := []models.CustomerPartnership{}
	db.PSQL.Find(&data)
	c.Set("data", data)
	c.Next()
}

func CreateCustomer(c *gin.Context) {

	userCtx, _ := c.Get("user")
	user := userCtx.(jwt.MapClaims)

	userId := user["sub"].(string)

	var body CreateCustomerBody

	c.ShouldBindBodyWith(&body, binding.JSON)

	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	var data models.Customer = models.Customer{
		Name:                       body.Name,
		BusinessRegistrationNumber: &body.BusinessRegistrationNumber,
		Url:                        &body.Url,
		Remarks:                    &body.Remarks,
		CustomerTypeId:             body.CustomerTypeId,
		CustomerPartnershipId:      body.CustomerPartnershipId,
		CountryId:                  body.CountryId,
		PartyId:                    body.PartyId,
		FileId:                     body.FileId,
		CreatedByUserId:            userId,
	}

	db.PSQL.Clauses(clause.Returning{}).Create(&data)
	db.PSQL.Preload("CustomerType").Preload("CustomerPartnership").Preload("Country").Preload("Party").Preload("Party.Country").Preload("Party.File").Preload("File").Find(&data)

	c.Set("data", data)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	partyId := c.Query("party_id")

	data, exist := customer_repository.GetCustomerByIdAndPartyId(id, partyId)

	if !exist {
		helpers.ThrowNotFoundError(c, fmt.Sprintf("Customer with id %s not found", id))
		return
	}

	var body UpdateCustomerBody

	c.ShouldBindBodyWith(&body, binding.JSON)

	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	data.Name = body.Name
	data.BusinessRegistrationNumber = &body.BusinessRegistrationNumber
	data.Url = &body.Url
	data.Remarks = &body.Remarks
	data.CustomerTypeId = body.CustomerTypeId
	data.CustomerPartnershipId = body.CustomerPartnershipId
	data.CountryId = body.CountryId
	data.FileId = body.FileId

	db.PSQL.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&data)
	db.PSQL.Preload("CustomerType").Preload("CustomerPartnership").Preload("Country").Preload("Party").Preload("Party.Country").Preload("Party.File").Preload("File").Find(&data)

	c.Set("data", data)
}

func CreateCustomerAddress(c *gin.Context) {
	customerId := c.Param("id")
	partyId := c.Query("party_id")

	var body CustomerAddressesBodyPayload

	c.ShouldBindBodyWith(&body, binding.JSON)

	_, exist := customer_repository.GetCustomerByIdAndPartyId(customerId, partyId)

	if !exist {
		helpers.ThrowNotFoundError(c, fmt.Sprintf("Customer with id %s not found", customerId))
		return
	}

	db.PSQL.Transaction(func(tx *gorm.DB) error {

		for index, address := range body.Addresses {
			err := initializers.Validate.Struct(address)

			if err != nil {
				c.JSON(400, gin.H{"error": err.Error() + " at index " + fmt.Sprint(index)})
				c.Abort()
				return errors.New("validation error")
			}

			_, exist := countries.FindById(address.CountryId)

			if !exist {
				helpers.ThrowNotFoundError(c, fmt.Sprintf("Country with id %s not found (at index %d)", address.CountryId, index))
				return errors.New("country not found")
			}

			var data models.CustomerAddresses = models.CustomerAddresses{
				AddressLine1: address.AddressLine1,
				AddressLine2: address.AddressLine2,
				AddressLine3: &address.AddressLine3,
				PostalCode:   address.PostalCode,
				CountryId:    address.CountryId,
				CustomerId:   customerId,
			}

			tx.Clauses(clause.Returning{}).Create(&data)
		}
		return nil
	})

	c.Set("data", body.Addresses)

}

func CreateContacts(c *gin.Context) {
	customerId := c.Param("id")
	partyId := c.Query("party_id")

	var body CustomerContactsBodyPayload

	c.ShouldBindBodyWith(&body, binding.JSON)

	_, exist := customer_repository.GetCustomerByIdAndPartyId(customerId, partyId)

	if !exist {
		helpers.ThrowNotFoundError(c, fmt.Sprintf("Customer with id %s not found", customerId))
		return
	}

	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		c.Next()
		return
	}

	db.PSQL.Transaction(func(tx *gorm.DB) error {

		for _, contact := range body.Contacts {

			var data models.CustomerContact = models.CustomerContact{
				Email:       contact.Email,
				Name:        contact.Name,
				PhoneNumber: contact.PhoneNumber,
				CustomerId:  customerId,
			}

			tx.Create(&data)
		}

		return nil
	})

	c.Set("data", body)
}
