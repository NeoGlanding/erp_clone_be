package controllers

import (
	"fmt"

	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm/clause"
)

type BodyPostParty struct {
	CompanyName		string		`validate:"required" json:"company_name"`
	AddressLine1	string		`validate:"required" json:"address_line_1"`
	AddressLine2	string		`json:"address_line_2"`
	AddressLine3	string		`json:"address_line_3"`
	PostalCode		string		`json:"postal_code"`
	CountryId		string		`json:"country_id"`
}

func GetParties(c *gin.Context) {

	var data []models.UserPartyPermission

	x, _ := c.Get("user")
	user := x.(jwt.MapClaims)

	db.PSQL.Limit(10).Where("user_id = ?", user["sub"]).Preload("Party").Preload("User").Preload("Party.Country").Find(&data)

	var formattedData []models.Party = []models.Party{}

	for _, item := range data {
		var singleData models.Party = models.Party{}
		singleData = item.Party
		add :=append(formattedData, singleData)
		formattedData = add;
	}

	fmt.Println(formattedData)

	c.Set("data", formattedData)
}

func PostParty(c *gin.Context) {
	x, _ := c.Get("user")
	user := x.(jwt.MapClaims)

	var body BodyPostParty;
	c.BindJSON(&body)

	err := initializers.Validate.Struct(body)

	if err != nil {
		err := helpers.DestructValidationError(&err)
		c.Set("error", err)
		c.Set("error-type", constants.REQUEST_VALIDATION_ERROR)
		c.Next()
	} 

	party := models.Party{
		Name: body.CompanyName,
		AddressLine1: body.AddressLine1,
		AddressLine2: &body.AddressLine2,
		AddressLine3: &body.AddressLine3,
		PostalCode: body.PostalCode,
		CountryId: body.CountryId,
	}

	resultParty := db.PSQL.Clauses(clause.Returning{}).Preload("Country").Create(&party)

	permission := models.UserPartyPermission{
		UserId: user["sub"].(string),
		PartyId: party.ID,
		Permission: "ADMIN",
	}

	resultPermission := db.PSQL.Create(&permission)

	errorCountryId := "ERROR: insert or update on table \"parties\" violates foreign key constraint \"fk_parties_country\" (SQLSTATE 23503)"

	if (resultParty.Error != nil) {
		if (resultParty.Error.Error() == errorCountryId) {
			c.Set("error", "Country ID is not found")
			c.Next()
		}
	}

	if (resultPermission.Error != nil) {
		fmt.Println(resultPermission.Error.Error())
	}


	c.Set("data", map[string]interface{}{"message": "Successfuly create party", "data": party})
	

}