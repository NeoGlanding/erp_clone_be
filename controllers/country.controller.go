package controllers

import (
	"github.com/automa8e_clone/helpers"
	"github.com/automa8e_clone/repositories/countries"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	data, _ := countries.FindAll()

	c.Set("data", data)
}

func GetCountry(c *gin.Context) {

	id := c.Param("id")

	data, exist := countries.FindById(id); if !exist {
		helpers.SetNotFoundError(c, "Country ID is not found")
		return
	}

	c.Set("data", data)
}