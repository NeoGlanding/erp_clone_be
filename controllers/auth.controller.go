package controllers

import (
	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

type RegisterBody struct {
	Email 		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required"`
	Phone		string	`json:"phone" validate:"required,e164"`
}

func Login(c *gin.Context) {
	var body LoginBody;
	c.BindJSON(&body)
	c.JSON(200, gin.H{"email": body.Email, "password": body.Password})
}

func Register(c *gin.Context) {
	var body RegisterBody;
	c.BindJSON(&body)
	
	err := initializers.Validate.Struct(body)

	// Hash
	hashed, _ :=bcrypt.GenerateFromPassword([]byte(body.Password), 8)

	if (err != nil) {
		c.Set("error", helpers.DestructValidationError(&err))
		c.Set("error-type", constants.REQUEST_VALIDATION_ERROR)
		return
	}

	user := models.User{
		Email: body.Email,
		Password: string(hashed),
		Phone: &body.Phone,
	}

	var existingUser models.User

	result := db.PSQL.Where("email = ? OR phone = ?", body.Email, body.Phone).First(&existingUser)

	if (result.RowsAffected > 0) {
		c.Set("error", "User already exist")
	}

	db.PSQL.Create(&user)

	c.Set("data", map[string]interface{}{"message": "Account successfuly created!"})

}