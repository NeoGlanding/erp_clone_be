package controllers

import (
	"errors"
	"fmt"
	"time"

	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/middlewares"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	userpartypermissions "github.com/automa8e_clone/repositories/user-party-permissions"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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
	Permission		string			`json:"permission"`
}

type PartyUserElement struct {
	ID				string			`json:"id"`
	FirstName		string			`json:"first_name"`
	Surname			string			`json:"surname"`
	Email			string			`json:"email"`
	PostalCode		string			`json:"postal_code"`
	Country			string			`json:"country"`
	Permission		string			`json:"permission"`
}
type ReturnParties struct {
	Data       []PartyListElement 	`json:"data"`
	Pagination types.PaginationResponse `json:"pagination"`
}

type ReturnPartyUsers struct {
	Data       []PartyUserElement 	`json:"data"`
	Pagination types.PaginationResponse `json:"pagination"`
}

type BodyPostParty struct {
	CompanyName  string `validate:"required" json:"company_name"`
	AddressLine1 string `validate:"required" json:"address_line_1"`
	AddressLine2 string `json:"address_line_2"`
	AddressLine3 string `json:"address_line_3"`
	PostalCode   string `json:"postal_code"`
	CountryId    string `json:"country_id"`
	FileId		 *string`json:"file_id,omitempty" validate:"omitempty,uuid"`
}

type BodyPostAction struct {
	PartyId		string		`json:"party_id" validate:"required,uuid"`
	Action		string		`json:"action" validate:"oneof=revoke viewer revoke_admin admin"`
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
	Table("user_party_permissions AS upp").
	Unscoped().
	Joins("JOIN parties ON parties.id = upp.party_id").
	Joins("JOIN countries ON parties.country_id = countries.id").
	Select("parties.id, parties.name, parties.created_at, parties. postal_code, parties.address_line1, countries.name AS country, upp.permission").
	Where("upp.user_id = ?", user["sub"]).
	Where("upp.deleted_at IS NULL")


	
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
	c.ShouldBindBodyWith(&body, binding.JSON)

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
		FileId: body.FileId,
	}

	var partyReturn models.Party


	resultParty := db.PSQL.Clauses(clause.Returning{}).Preload("Country").Create(&party)
	db.PSQL.Preload("Country").Preload("File").Where("id = ?", party.ID).Find(&partyReturn)

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

	c.Set("data", map[string]interface{}{"message": "Successfuly create party", "data": partyReturn})

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

	c.ShouldBindBodyWith(&body, binding.JSON)

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
		FileId: body.FileId,
	}

	db.PSQL.Model(&models.Party{}).Where("id = ?", id).Save(&args)

	db.PSQL.Preload("Country").Preload("File").Where("id = ?", id).Find(&data)

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

	db.PSQL.Preload("File").Preload("Country").Where("id = ?", id).Find(&party)

	c.Set("data", party)
}

func PartyAction(c *gin.Context) {
	var body BodyPostAction;

	userCtx, _ := c.Get("user"); myUser := userCtx.(jwt.MapClaims)

	c.ShouldBindBodyWith(&body, binding.JSON)

	fmt.Println(body)

	err := initializers.Validate.Struct(body);


	// if body.Action != "viewer" {
	// 	helpers.SetForbiddenError(c, "Forbidden resources")
	// 	return
	// }

	if err != nil {
		fmt.Println("are you supposed to be cop or something", err)
		helpers.SetValidationError(c, &err)
	}

	err = db.PSQL.Transaction(func(tx *gorm.DB) error {

		
		for _, email := range body.UserEmails {
			var user models.User
			err := tx.Where("email = ?", email).First(&user).Error

			if err != nil {
				return err
			}

			_, isOnboarded := users_repository.CheckIsOnboarded(user.Id)

			if !isOnboarded {
				message := fmt.Sprintf("failed: %s is not onboarded", user.Email)
				return errors.New(message)
			}

			permission, _ := userpartypermissions.RetrievePermission(body.PartyId, myUser["sub"].(string))
			isOwner := permission == types.PERMISSION_OWNER

			fmt.Println(permission)

			value := models.UserPartyPermission{
				UserId: user.Id,
				PartyId: body.PartyId,
				Permission: "VIEWER",
			}

			if (body.Action == "admin") {
				value.Permission = "ADMIN"
			}

			var query *gorm.DB
			
			if body.Action == "viewer" {
				query = tx.Table("user_party_permissions").Where("user_id = ? AND party_id = ? AND permission = 'VIEWER'", value.UserId, value.PartyId).FirstOrCreate(&value)
			} else if body.Action == "revoke" {
				query = tx.Where("user_id = ? AND party_id = ? AND permission = 'VIEWER'", user.Id, body.PartyId).Delete(&value)
			} else if body.Action == "admin" {
				if !isOwner {
					helpers.SetForbiddenError(c, "Forbidden Resources")
					return errors.New("you are not owner of this party")
				} else {
					query = tx.Table("user_party_permissions").Where("user_id = ? AND party_id = ? AND permission = 'ADMIN'", value.UserId, value.PartyId).FirstOrCreate(&value)
				}
			} else if body.Action == "revoke_admin" {
				if !isOwner {
					helpers.SetForbiddenError(c, "Forbidden Resources")
					return errors.New("you are not owner of this party")
				} else {
					query = tx.Where("user_id = ? AND party_id = ? AND permission = 'ADMIN'", user.Id, body.PartyId).Delete(&value)
				}
			}

			fmt.Println(query.RowsAffected)
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

func GetPartyUsers(c *gin.Context) {
	id := c.Param("id");
	paginationCtx, _ := c.Get("pagination"); pagination := paginationCtx.(types.PaginationQuery)
	queryCtx, _ := c.Get("query"); query := queryCtx.(middlewares.TypeQueryMiddleware)

	var data []PartyUserElement;
	var count int64;

	fmt.Println("pagination -> ", pagination)
	fmt.Println("query -> ", query)

	base := db.PSQL.Table("user_party_permissions AS upp").
			Joins("JOIN users ON users.id = upp.user_id").
			Joins("JOIN user_details AS details ON details.user_id = users.id").
			Joins("JOIN countries ON countries.id = details.country_id").
			Select("users.id, details.first_name, details.surname, users.email, details.postal_code, countries.name AS country, upp.permission").
			Order("CASE upp.permission WHEN 'OWNER' THEN 1 WHEN 'ADMIN' THEN 2 WHEN 'VIEWER' THEN 3 ELSE 4 END").
			Where("upp.party_id = ?", id)

	if query.SearchExist {
		base = base.Where("LOWER(details.first_name) LIKE LOWER(?) OR LOWER(details.surname) LIKE LOWER(?) OR LOWER(users.email) LIKE LOWER(?)", query.Search, query.Search, query.Search)
	}

	base.
		Offset(pagination.Offset).
		Limit(pagination.PageSize).
		Count(&count).
		Find(&data)

	response := ReturnPartyUsers{
		Data: data,
		Pagination: types.PaginationResponse{
			TotalData: count,
			TotalPage: int(helpers.FindTotalPage(count, pagination.PageSize)),
			CurrentPage: pagination.Page,
			PageSize: pagination.PageSize,
		},
	}

	c.Set("data", response)

}