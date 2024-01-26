package middlewares

import "github.com/gin-gonic/gin"

func InterceptParam(param string, properties string) (func (c *gin.Context)) {
	return func (c *gin.Context) {
		value := c.Param(param)
		c.Set(properties, value)
	}
}