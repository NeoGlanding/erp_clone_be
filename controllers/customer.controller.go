package controllers

import (
	"github.com/automa8e_clone/db"
	"github.com/automa8e_clone/models"
	"github.com/gin-gonic/gin"
)

func GetCustomerType(c *gin.Context) {
	data := []models.CustomerType{}
	db.PSQL.Find(&data)
	c.Set("data", data)
	c.Next()
}

func GetCustomerPartnership(c *gin.Context) {
	data := []models.CustomerPartnership{}
	db.PSQL.Find(&data)
	c.Set("data", data)
	c.Next()
}