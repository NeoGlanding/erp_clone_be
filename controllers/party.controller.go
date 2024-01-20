package controllers

import "github.com/gin-gonic/gin"

func GetParty(c *gin.Context) {
	c.JSON(200, map[string]interface{}{"message": "success go here"})
}