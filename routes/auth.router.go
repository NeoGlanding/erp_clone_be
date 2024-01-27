package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/login", controllers.Login, middlewares.ResponseMiddlewares)
	auth.POST("/register", controllers.Register, middlewares.ResponseMiddlewares)
	auth.POST("/refresh-token", controllers.RetrieveAccessToken, middlewares.ResponseMiddlewares)
}