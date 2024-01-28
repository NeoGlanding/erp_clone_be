package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostFile(c *gin.Context) {

	userCtx, _ := c.Get("user"); user := userCtx.(jwt.MapClaims)

	file, _ := c.FormFile("file")
	log.Println(file.Filename)

	c.Set("data", user)
}