package middlewares

import (
	"fmt"
	"net/http"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	"github.com/automa8e_clone/repositories/files"
	userpartypermissions "github.com/automa8e_clone/repositories/user-party-permissions"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
)

type BodyEmail struct {
	Email string `json:"email"`
}

type BodyCountryId struct {
	CountryId string `json:"country_id"`
}

type BodyPartyId struct {
	PartyId string `json:"party_id"`
}

type BodyCustomerPartnershipId struct {
	CustomerPartnershipId string `json:"customer_partnership_id"`
}

type BodyCustomerTypeId struct {
	CustomerTypeId string `json:"customer_type_id"`
}

func BodyEmailExistMiddleware(c *gin.Context) {
	var body BodyEmail
	c.ShouldBindBodyWith(&body, binding.JSON)
	var user models.User
	exist := users_repository.FindByEmail(body.Email, &user)

	if !exist {
		helpers.ThrowError(c, http.StatusBadRequest, "Email is not found")
		c.Abort()
		return
	}
}

func BodyCountryIdExistMiddleware(c *gin.Context) {
	var body BodyCountryId
	c.ShouldBindBodyWith(&body, binding.JSON)
	_, exist := countries.FindById(body.CountryId)

	fmt.Println("validation middleware body country_id", exist)

	if !exist {
		helpers.ThrowError(c, http.StatusBadRequest, "Invalid Country ID")
		return
	}

	c.Next()
}

func FileIdExist(c *gin.Context) {
	fileIdCtx, _ := c.Get("file-id")
	fileId := fileIdCtx.(string)
	userCtx, _ := c.Get("user")
	user := userCtx.(jwt.MapClaims)

	if fileId == "" {
		c.Next()
		return
	}

	_, exist := files.GetFileById(fileId, user["sub"].(string))

	if !exist {
		helpers.ThrowError(c, http.StatusBadRequest, "File not found")
		return
	}

	c.Next()
}

func PartyIdExistBody(c *gin.Context) {
	var userCtx, _ = c.Get("user")
	user := userCtx.(jwt.MapClaims)
	var body BodyPartyId
	c.ShouldBindBodyWith(&body, binding.JSON)
	_, exist := userpartypermissions.RetrievePermission(body.PartyId, user["sub"].(string))
	if !exist {
		helpers.ThrowError(c, http.StatusBadRequest, "Party ID not found")
		return
	}
	c.Next()
}

func CustomerPartnershipIdExistBody(c *gin.Context) {
	var body BodyCustomerPartnershipId
	c.ShouldBindBodyWith(&body, binding.JSON)

	var data models.CustomerPartnership

	tx := db.PSQL.Table("customer_partnerships").Where("id = ?", body.CustomerPartnershipId).Find(&data)

	if tx.RowsAffected == 0 {
		helpers.ThrowError(c, http.StatusBadRequest, "Customer Partnership ID not found")
		return
	}

	c.Next()
}

func CustomerTypeIdExistBody(c *gin.Context) {
	var body BodyCustomerTypeId
	c.ShouldBindBodyWith(&body, binding.JSON)

	var data models.CustomerType

	fmt.Println("body.CustomerTypeId", body.CustomerTypeId)

	tx := db.PSQL.Table("customer_types").Where("id = ?", body.CustomerTypeId).Find(&data)

	if tx.RowsAffected == 0 {
		helpers.ThrowError(c, http.StatusBadRequest, "Customer Types ID not found")
		return
	}

	c.Next()
}
