package middlewares

import (
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BodyEmail struct {
	Email	string	`json:"email"`
}

type BodyCountryId struct {
	CountryId	string	`json:"country_id"`
}

func BodyEmailExistMiddleware(c *gin.Context) {
	var body BodyEmail
	c.ShouldBindBodyWith(&body, binding.JSON)
	var user models.User
	exist := users_repository.FindByEmail(body.Email, &user)

	if !exist {
		helpers.SetBadRequestError(c, "Email is not found")
		return
	}
}

func BodyCountryIdExistMiddleware(c *gin.Context) {
	var body BodyCountryId
	c.ShouldBindBodyWith(&body, binding.JSON)
	_,exist := countries.FindById(body.CountryId)

	if !exist {
		helpers.SetBadRequestError(c, "Country ID is not found")
		return
	}
}