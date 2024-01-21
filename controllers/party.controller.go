package controllers

import (
	"fmt"
	"time"

	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

type PartyListElement struct {
	ID				string			`json:"id"`
	Name			string			`json:"name"`
	Country			string			`json:"country"`
	AddressLine1	string			`json:"address_line_1"`
	CreatedAt		time.Time		`json:"created_at"`
}
type ReturnParties struct {
	Data       []PartyListElement `json:"data"`
	Pagination types.PaginationResponse `json:"pagination"`
}

type BodyPostParty struct {
	CompanyName  string `validate:"required" json:"company_name"`
	AddressLine1 string `validate:"required" json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	PostalCode   string `json:"postal_code"`
	CountryId    string `json:"country_id"`
}

func GetParties(c *gin.Context) {

	var data []PartyListElement
	var count int64

	userCtx, _ := c.Get("user")
	user := userCtx.(jwt.MapClaims)

	paginationCtx, _ := c.Get("pagination")
	pagination := paginationCtx.(types.PaginationQuery)

	db.PSQL.
	Table("user_party_permissions").
	Offset(pagination.Offset).
	Limit(pagination.PageSize).
	Joins("JOIN parties ON parties.id = user_party_permissions.party_id").
	Joins("JOIN countries ON parties.country_id = countries.id").
	Select("parties.id, parties.name, parties.created_at, parties. postal_code, parties.address_line1, countries.name AS country").
	Order("parties.created_at desc").
	Where("user_party_permissions.user_id = ?", user["sub"]).
	Find(&data).
	Count(&count)

	// Find Total Data
	// db.PSQL.Model(&models.UserPartyPermission{}).Where("user_id = ?", user["sub"]).Count(&count)

	response := ReturnParties{
		Data: data,
		Pagination: types.PaginationResponse{
			TotalData:   count,
			CurrentPage: pagination.Page,
			TotalPage:   int(helpers.FindTotalPage(count, pagination.PageSize)),
			PageSize:    pagination.PageSize,
		},
	}

	c.Set("data", response)
}

func PostParty(c *gin.Context) {
	x, _ := c.Get("user")
	user := x.(jwt.MapClaims)

	var body BodyPostParty
	c.BindJSON(&body)

	err := initializers.Validate.Struct(body)

	if err != nil {
		err := helpers.DestructValidationError(&err)
		c.Set("error", err)
		c.Set("error-type", constants.REQUEST_VALIDATION_ERROR)
		c.Next()
	}

	party := models.Party{
		Name:         body.CompanyName,
		AddressLine1: body.AddressLine1,
		AddressLine2: &body.AddressLine2,
		AddressLine3: &body.AddressLine3,
		PostalCode:   body.PostalCode,
		CountryId:    body.CountryId,
	}

	resultParty := db.PSQL.Clauses(clause.Returning{}).Preload("Country").Create(&party)

	permission := models.UserPartyPermission{
		UserId:     user["sub"].(string),
		PartyId:    party.ID,
		Permission: "ADMIN",
	}

	resultPermission := db.PSQL.Create(&permission)

	errorCountryId := "ERROR: insert or update on table \"parties\" violates foreign key constraint \"fk_parties_country\" (SQLSTATE 23503)"

	if resultParty.Error != nil {
		if resultParty.Error.Error() == errorCountryId {
			c.Set("error", "Country ID is not found")
			c.Next()
		}
	}

	if resultPermission.Error != nil {
		fmt.Println(resultPermission.Error.Error())
	}

	c.Set("data", map[string]interface{}{"message": "Successfuly create party", "data": party})

}
