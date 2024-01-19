package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseMiddlewares(c *gin.Context) {
	data, _ := c.Get("data")

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})

}