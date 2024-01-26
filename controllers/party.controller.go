package controllers

import (
	"fmt"
	"time"

	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/middlewares"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
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

type BodyPostAction struct {
	PartyId		string		`json:"party_id" validate:"required,uuid"`
	Action		string		`json:"action" validate:"oneof=revoke viewer"`
	UserEmails	[]string	`json:"user_emails" validate:"required,min=1,dive,email"`
}

func GetParties(c *gin.Context) {

	var data []PartyListElement
	var count int64

	userCtx, _ := c.Get("user")
	user := userCtx.(jwt.MapClaims)

	paginationCtx, _ := c.Get("pagination")
	pagination := paginationCtx.(types.PaginationQuery)

	queryCtx, _ := c.Get("query")
	query := queryCtx.(middlewares.TypeQueryMiddleware)


	base := db.PSQL.
	Table("user_party_permissions").
	Joins("JOIN parties ON parties.id = user_party_permissions.party_id").
	Joins("JOIN countries ON parties.country_id = countries.id").
	Select("parties.id, parties.name, parties.created_at, parties. postal_code, parties.address_line1, countries.name AS country").
	Where("user_party_permissions.user_id = ?", user["sub"])


	
	if query.SearchExist{
		base = base.Where("LOWER(parties.name) LIKE LOWER(?) OR Lower(parties.address_line1) LIKE LOWER(?) OR Lower(parties.address_line2) LIKE LOWER(?) OR Lower(parties.address_line3) LIKE LOWER(?)", query.Search, query.Search, query.Search, query.Search)
	}
	
	if query.SortByExist {
		base = base.Order(fmt.Sprintf("%s %s", query.SortBy, query.SortDirection))
	}
	
	base.
	Count(&count).
	Offset(pagination.Offset).
	Limit(pagination.PageSize).
	Find(&data)


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
		Permission: "OWNER",
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

func UpdateParty(c *gin.Context) {
	id := c.Param("id");
	userCtx, _ := c.Get("user")

	user := userCtx.(jwt.MapClaims)

	var permission models.UserPartyPermission
	var body BodyPostParty

	query := db.PSQL.Table("user_party_permissions").
		Clauses(clause.Locking{Strength: "UPDATE"}).
		Where("party_id = ? AND user_id = ?", id, user["sub"]).
		Find(&permission)

	if query.RowsAffected == 0 {
		helpers.SetNotFoundError(c, "Party not found")
		return
	}

	c.BindJSON(&body)

	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	_, exist := countries.FindById(body.CountryId); if !exist {
		helpers.SetBadRequestError(c, "Country ID is not found")
		return
	}

	var data models.Party
	var args models.Party = models.Party{
		ID: id,
		Name: body.CompanyName,
		AddressLine1: body.AddressLine1,
		AddressLine2: &body.AddressLine2,
		AddressLine3: &body.AddressLine3,
		PostalCode: body.PostalCode,
		CountryId: body.CountryId,
	}

	db.PSQL.Model(&models.Party{}).Where("id = ?", id).Save(&args)

	db.PSQL.Preload("Country").Where("id = ?", id).Find(&data)

	c.Set("data", map[string]interface{}{"message": "Success", "data": data})
}

func GetParty(c *gin.Context) {
	id := c.Param("id")

	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	var data models.UserPartyPermission

	query := db.PSQL.Where("user_id = ? AND party_id = ?", user["sub"], id).Find(&data);

	if query.RowsAffected == 0 {
		helpers.SetNotFoundError(c,"Party not found")
		return
	}

	var party models.Party;

	db.PSQL.Preload("Country").Where("id = ?", id).Find(&party)

	c.Set("data", party)
}

func PartyAction(c *gin.Context) {
	var body BodyPostAction;

	c.BindJSON(&body)

	err := initializers.Validate.Struct(body);

	if body.Action != "viewer" {
		c.JSON(400,gin.H{"message": "232i9"})
		c.Abort()
	}

	if err != nil {
		helpers.SetValidationError(c, &err)
	}

	err = db.PSQL.Transaction(func(tx *gorm.DB) error {

		
		for _, email := range body.UserEmails {
			var user models.User
			err := tx.Where("email = ?", email).First(&user).Error

			if err != nil {
				
				return err
			}

			value := models.UserPartyPermission{
				UserId: user.Id,
				PartyId: body.PartyId,
				Permission: "VIEWER",
			}

			var query *gorm.DB
			
			if body.Action == "viewer" {
				query = tx.Table("user_party_permissions").Where("user_id = ? AND party_id = ? AND permission = 'VIEWER'", value.UserId, value.PartyId).FirstOrCreate(&value)
			}

		}

		return nil
	})

	if err != nil {
		helpers.SetBadRequestError(c,err.Error())
		c.Next()
	} else {
		c.Set("data", "Successfully modify access")

	}

}