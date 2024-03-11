package controllers

import (
	"fmt"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	customer_repository "github.com/automa8e_clone/repositories/customers"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
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

	c.Set("data", data)
}
