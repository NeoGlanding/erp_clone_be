package controllers

import (
	"errors"
	"fmt"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/middlewares"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	customer_repository "github.com/automa8e_clone/repositories/customers"
	"github.com/automa8e_clone/types"
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

type UpdateCustomerAddressesBody struct {
	ID           string `json:"id" validate:"required,uuid"`
	AddressLine1 string `json:"address_line1" validate:"required"`
	AddressLine2 string `json:"address_line2" validate:"required"`
	AddressLine3 string `json:"address_line3"`
	PostalCode   string `json:"postal_code" validate:"required"`
	CountryId    string `json:"country_id" validate:"required,uuid"`
}

type UpdateCustomerAddressesBodyPayload struct {
	Addresses []UpdateCustomerAddressesBody `json:"addresses" validate:"required"`
}

type UpdateCustomerContactsBody struct {
	ID          string `json:"id" validate:"required,uuid"`
	Name        string `json:"name" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phone_number" validate:"required,e164"`
}

type UpdateCustomerContactBodyPayload struct {
	Contacts []UpdateCustomerContactsBody `json:"contacts" validate:"required,dive,required"`
}

type ReturnCustomers struct {
	Data       []models.Customer        `json:"data"`
	Pagination types.PaginationResponse `json:"pagination"`
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

	userCtx, _ := c.Get("user")
	user := userCtx.(jwt.MapClaims)

	data, exist := customer_repository.GetCustomerByIdAndPartyId(id, partyId)

	var arg models.Customer

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

	arg.ID = id
	arg.Name = body.Name
	arg.BusinessRegistrationNumber = &body.BusinessRegistrationNumber
	arg.Url = &body.Url
	arg.Remarks = &body.Remarks
	arg.CustomerTypeId = body.CustomerTypeId
	arg.CustomerPartnershipId = body.CustomerPartnershipId
	arg.CountryId = body.CountryId
	arg.FileId = body.FileId
	arg.PartyId = partyId
	arg.CreatedByUserId = user["sub"].(string)

	db.PSQL.Clauses(clause.Locking{Strength: "UPDATE"}).Save(&arg)
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
				c.Abort()
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

func UpdateCustomerAddresses(c *gin.Context) {
	customerId := c.Param("id")
	partyId := c.Query("party_id")

	var body UpdateCustomerAddressesBodyPayload

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

			_, addressExist := customer_repository.GetCustomerAddressById(address.ID, customerId)

			if !addressExist {
				helpers.ThrowNotFoundError(c, fmt.Sprintf("Address with id %s not found (at index %d)", address.ID, index))
				c.Abort()
				return errors.New("address not found")
			}

			_, exist := countries.FindById(address.CountryId)

			if !exist {
				helpers.ThrowNotFoundError(c, fmt.Sprintf("Country with id %s not found (at index %d)", address.CountryId, index))
				c.Abort()
				return errors.New("country not found")
			}

			var data models.CustomerAddresses = models.CustomerAddresses{
				ID:           address.ID,
				AddressLine1: address.AddressLine1,
				AddressLine2: address.AddressLine2,
				AddressLine3: &address.AddressLine3,
				PostalCode:   address.PostalCode,
				CountryId:    address.CountryId,
				CustomerId:   customerId,
			}

			tx.Clauses(clause.Returning{}).Save(&data)
		}
		return nil
	})

	c.Set("data", body.Addresses)
}

func UpdateContacts(c *gin.Context) {
	customerId := c.Param("id")
	partyId := c.Query("party_id")

	var body UpdateCustomerContactBodyPayload

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

			_, exist := customer_repository.GetCustomerContactById(contact.ID, customerId)

			if !exist {
				helpers.ThrowBadRequestError(c, fmt.Sprintf("Contact ID %s is not exist", contact.ID))
			}

			var data models.CustomerContact = models.CustomerContact{
				ID:          contact.ID,
				Email:       contact.Email,
				Name:        contact.Name,
				PhoneNumber: contact.PhoneNumber,
				CustomerId:  customerId,
			}

			tx.Save(&data)
		}

		return nil
	})

	c.Set("data", body)
}

func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	partyId := c.Query("party_id")

	data, exist := customer_repository.GetCustomerByIdAndPartyId(id, partyId)

	if !exist {
		helpers.ThrowNotFoundError(c, "Customer not found")
		return
	}

	c.Set("data", data)

}

func GetCustomers(c *gin.Context) {

	paginationCtx, _ := c.Get("pagination")
	pagination := paginationCtx.(types.PaginationQuery)

	partyId := c.Query("party_id")

	queryCtx, _ := c.Get("query")
	query := queryCtx.(middlewares.TypeQueryMiddleware)

	var data []models.Customer
	var count int64

	base := db.PSQL.
		Table("customers").
		Where("party_id = ?", partyId)

	if query.SearchExist {
		base = base.Where("LOWER(name) LIKE LOWER(?)", query.Search)
	}

	if query.SortByExist {
		base = base.Order(fmt.Sprintf("%s %s", query.SortBy, query.SortDirection))
	}

	base.Count(&count).
		Preload("Country").
		Preload("Party").
		Preload("File").
		Offset(pagination.Offset).Limit(pagination.PageSize).Find(&data)

	c.Set("data", ReturnCustomers{
		Data: data,
		Pagination: types.PaginationResponse{
			TotalData:   count,
			TotalPage:   helpers.FindTotalPage(count, pagination.PageSize),
			CurrentPage: pagination.Page,
			PageSize:    pagination.PageSize,
		},
	})

}
