package middlewares

import (
	"fmt"
	"net/http"

	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/models"
	"github.com/automa8e_clone/repositories/countries"
	"github.com/automa8e_clone/repositories/files"
	users_repository "github.com/automa8e_clone/repositories/users"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
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
		helpers.ThrowError(c, http.StatusBadRequest, "Email is not found")
		c.Abort()
		return
	}
}

func BodyCountryIdExistMiddleware(c *gin.Context) {
	var body BodyCountryId
	c.ShouldBindBodyWith(&body, binding.JSON)
	_,exist := countries.FindById(body.CountryId)

	fmt.Println("data", exist)


	if !exist {
		helpers.ThrowError(c, http.StatusBadRequest, "Invalid Country ID")
		return
	}
}

func FileIdExist(c *gin.Context) {
	fileIdCtx, _ := c.Get("file-id"); fileId := fileIdCtx.(string)
	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)



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

