package controllers

import (
	"fmt"
	"strconv"

	"github.com/automa8e_clone/config"
	"github.com/automa8e_clone/constants"
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/initializers"
	"github.com/automa8e_clone/models"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email 		string 	`json:"email" validate:"required,email"`
	Password 	string 	`json:"password" validate:"required"`
}

type RegisterBody struct {
	Email 		string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required"`
	Phone		string	`json:"phone" validate:"required,e164"`
}

type PostRefreshTokenBody struct {
	RefreshToken	string	`json:"refresh_token" validate:"required,jwt"`
}

func Login(c *gin.Context) {
	var body LoginBody;
	c.BindJSON(&body)
	
	err := initializers.Validate.Struct(body)

	if (err != nil) {
		c.Set("error", helpers.DestructValidationError(&err))
		c.Set("error-type", constants.REQUEST_VALIDATION_ERROR)
		return
	}

	user := models.User{
		Email: body.Email,
	}

	db.PSQL.Where("email = ?", body.Email).Find(&user)


	// Compare hash
	passwordError := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	
	if (passwordError != nil) {
		c.Set("error", "Invalid credential")
		c.Set("error-code", 401)
		return
	}
	
	expireTime, err := strconv.Atoi(config.AppConfig.JWT_TIME_EXPIRATION)

	if err != nil {
		panic(err)
	}

	access_token, err := helpers.GenerateJWT(constants.ACCESS_TOKEN_REF,expireTime, &user);
	refresh_token, _ := helpers.GenerateJWT(constants.REFRESH_TOKEN_REF, 1440 * 7, &user)

	if err != nil {
		fmt.Println(err)
	}

	
	c.Set("data", map[string]interface{}{"access_token": access_token, "refresh_token": refresh_token})
}

func Register(c *gin.Context) {
	var body RegisterBody;
	c.BindJSON(&body)
	
	err := initializers.Validate.Struct(body)

	// Hash
	hashed, _ := bcrypt.GenerateFromPassword([]byte(body.Password), 8)

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

func RetrieveAccessToken(c *gin.Context) {
	var body PostRefreshTokenBody;
	var user models.User


	c.ShouldBindBodyWith(&body, binding.JSON)


	err := initializers.Validate.Struct(body)

	if err != nil {
		helpers.SetValidationError(c, &err)
		return
	}

	claims, err := helpers.ParseJWT(body.RefreshToken);
	users_repository.FindByEmail(claims["email"].(string), &user)

	if err != nil {
		helpers.SetBadRequestError(c, "Invalid refresh token")
		return
	}

	ref := claims["ref"].(string)
	ic := claims["ic"].(float64)


	if ref != constants.REFRESH_TOKEN_REF {
		helpers.SetBadRequestError(c, "Invalid refresh token")
		return
	}

	if ic != float64(user.InformationChanged) {
		helpers.SetBadRequestError(c, "Invalid refresh token")
		return
	}

	expireTime, err := strconv.Atoi(config.AppConfig.JWT_TIME_EXPIRATION)

	if err != nil {
		helpers.SetInternalServerError(c, "An error occured in system while generating new access token")
	}

	users_repository.FindByEmail(claims["email"].(string), &user)

	token, err := helpers.GenerateJWT(constants.ACCESS_TOKEN_REF, expireTime, &user)

	response := map[string]string{
		"access_token": token,
	}

	c.Set("data", response)
}