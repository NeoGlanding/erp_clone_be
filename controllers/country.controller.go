package controllers

import (
	"github.com/automa8e_clone/repositories/countries"
	"github.com/gin-gonic/gin"
)

func GetCountries(c *gin.Context) {
	data, _ := countries.FindAll()

	c.Set("data", data)
}