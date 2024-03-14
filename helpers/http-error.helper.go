package helpers

import (
	"net/http"

	"github.com/automa8e_clone/constants"
	"github.com/gin-gonic/gin"
)

func SetNotFoundError(c *gin.Context, message string) {
	c.Set("error", message)
	c.Set("error-code", http.StatusNotFound)
}

func SetBadRequestError(c *gin.Context, message string) {
	c.Set("error", message)
	c.Set("error-code", http.StatusBadRequest)
}

func SetForbiddenError(c *gin.Context, message string) {
	c.Set("error", message)
	c.Set("error-code", http.StatusForbidden)
}

func SetValidationError(c *gin.Context, err *error) {
	c.Set("error", DestructValidationError(err))
	c.Set("error-code", http.StatusBadRequest)
	c.Set("error-type", constants.REQUEST_VALIDATION_ERROR)
}

func SetInternalServerError(c *gin.Context, message string) {
	c.Set("error", message)
	c.Set("error-code", http.StatusInternalServerError)
}

func SetConflictError(c *gin.Context, message string) {
	c.Set("error", message)
	c.Set("error-code", http.StatusConflict)
}

func ThrowError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"message": message, "status": statusCode})
	c.Abort()
}

func ThrowBadRequestError(c *gin.Context, message string) {
	c.SecureJSON(http.StatusBadRequest, gin.H{"message": message, "status": http.StatusBadRequest})
	c.Abort()
}

func ThrowNotFoundError(c *gin.Context, message string) {
	c.SecureJSON(http.StatusNotFound, gin.H{"message": message, "status": http.StatusNotFound})
	c.Abort()
}
