package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Customer(c *gin.Engine) {
	customer := c.Group("/customers")

	customer.GET("/types", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,controllers.GetCustomerType, middlewares.ResponseMiddlewares)
	customer.GET("/partnerships", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization,controllers.GetCustomerPartnership, middlewares.ResponseMiddlewares)
}