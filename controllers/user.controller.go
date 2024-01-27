package controllers

import (
	"fmt"

	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type PostOnboardingBody struct {
	FirstName		string		`json:"first_name" validate:"required,max=100"`
	Surname			string		`json:"surname" validate:"required,max=100"`
	AddressLine1	string		`json:"address_line_1" validate:"required,max=100"`
	AddressLine2	*string		`json:"address_line_2" validate:"omitempty,max=100"`
	AddressLine3	*string		`json:"address_line_3" validate:"omitempty,max=100"`
	PostalCode		string		`json:"postal_code" validate:"required,max=100"`
	CountryId		string		`json:"country_id" validate:"required"`
	IdentityNumber	string		`json:"identity_number" validate:"required"`
	DateOfBirth		string		`json:"date_of_birth" validate:"required,datestring"`
}

type PutUpdateCredentials struct {
	Email		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,password"`
	PhoneNumber	string	`json:"phone_number" validate:"required,e164"`
}

func OnboardUser(c *gin.Context) {
	var body PostOnboardingBody;
	
	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	c.ShouldBindBodyWith(&body, binding.JSON)

	err := initializers.Validate.Struct(body);

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	var userDetails models.UserDetails

	date, _ := helpers.FormatToTimestamps(body.DateOfBirth)

	value := models.UserDetails {
		FirstName: body.FirstName,
		Surname: body.Surname,
		UserId: user["sub"].(string),
		AddressLine1: body.AddressLine1,
		AddressLine2: *body.AddressLine2,
		AddressLine3: *body.AddressLine3,
		PostalCode: body.PostalCode,
		CountryId: body.CountryId,
		IdentityNumber: body.IdentityNumber,
		DateOfBirth: date,
	}

	fmt.Println(value.UserId)

	db.PSQL.Table("user_details").Create(&value)
	db.PSQL.Table("user_details").Preload("Country").Preload("User").Where("user_id = ?", value.UserId).Find(&userDetails)

	c.Set("data", userDetails)
	
}

func UpdateCredential(c *gin.Context) {

	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	var body PutUpdateCredentials

	c.ShouldBindBodyWith(&body, binding.JSON)

	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	_, exist := users_repository.CheckIsExistByEmailAndPhone(body.Email, body.PhoneNumber);

	if exist {
		helpers.SetBadRequestError(c, "Account already exist")
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	var data models.User

	db.PSQL.Table("users").Where("id = ?", user["sub"]).Find(&data)

	data.Email = body.Email
	data.Password = string(hashedPassword)
	data.Phone = &body.PhoneNumber
	data.InformationChanged++

	db.PSQL.Save(&data)

	c.Set("data", "Successfuly updated")
}