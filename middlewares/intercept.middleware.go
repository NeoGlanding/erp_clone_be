package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type PartyIdBody struct {
	PartyId	 string `json:"party_id"`
}

type EmailBody struct {
	Email	string	`json:"email"`
}

type CountryIdBody struct {
	CountryId	string	`json:"country_id"`
}


func InterceptParam(param string, properties string) (func (c *gin.Context)) {
	return func (c *gin.Context) {
		value := c.Param(param)
		c.Set(properties, value)
	}
}

func InterceptPartyIdFromBody (c *gin.Context) {
	var body PartyIdBody
	c.ShouldBindBodyWith(&body, binding.JSON)
	c.Set("party-id", body.PartyId)
}

func InterceptEmailFromBody(c *gin.Context) {
	var	body EmailBody
	c.ShouldBindBodyWith(&body, binding.JSON)
	c.Set("email", body.Email)
}

func InterceptFileIdBody(field string) (func (c *gin.Context)) {
	return func (c *gin.Context) {
		var body map[string]string;

		c.ShouldBindBodyWith(&body, binding.JSON)

		c.Set("file-id", body[field])
	}
}

func InterceptCountryIdFromBody (c *gin.Context) {
	var body CountryIdBody
	c.ShouldBindBodyWith(&body, binding.JSON)
	c.Set("country-id", body.CountryId)
}