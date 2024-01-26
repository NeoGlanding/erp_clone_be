package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Country(r *gin.Engine) {
	country := r.Group("/countries")

	country.GET("/", middlewares.TokenAuthenticationMiddleware, controllers.GetCountries, middlewares.ResponseMiddlewares)
	country.GET("/:id", middlewares.TokenAuthenticationMiddleware, controllers.GetCountry, middlewares.ResponseMiddlewares)
}