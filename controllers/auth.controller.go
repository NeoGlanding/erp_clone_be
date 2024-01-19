package controllers

import "github.com/gin-gonic/gin"

type LoginBody struct {
	Email 		string 	`json:"email"`
	Password 	string 	`json:"password"`
}

func Login(c *gin.Context) {
	var body LoginBody;
	c.BindJSON(&body)
	c.JSON(200, gin.H{"email": body.Email, "password": body.Password})
}