package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("/onboard", middlewares.TokenAuthenticationMiddleware, middlewares.UnonboardedAuthorization, middlewares.BodyCountryIdExistMiddleware, controllers.OnboardUser,middlewares.ResponseMiddlewares)
}