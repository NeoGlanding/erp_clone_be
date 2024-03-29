package middlewares

import (
	"net/http"

	"github.com/automa8e_clone/constants"
	"github.com/gin-gonic/gin"
)



func ResponseMiddlewares(c *gin.Context) {
	data, _ := c.Get("data")
	err, _ := c.Get("error")
	errType, _ := c.Get("error-type")
	errCode, _ := c.Get("error-code")

	if (err != nil) {
		if (errType == constants.REQUEST_VALIDATION_ERROR) {
			c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "errors": err})
		} else {
			var defaultErrorCode int = 400
			if (errCode != nil) {
				defaultErrorCode = errCode.(int)
			}
			c.JSON(defaultErrorCode, gin.H{"status": defaultErrorCode, "message": err})
		}

		return
	}


	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": data})

}