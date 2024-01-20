package controllers

import "github.com/gin-gonic/gin"

func GetParty(c *gin.Context) {

	c.Set("data", map[string]interface{}{"message": "Success"})
}

func PostParty(c *gin.Context) {
	c.Set("data", "dsjkj")
}