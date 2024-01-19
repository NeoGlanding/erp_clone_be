package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register, middlewares.ResponseMiddlewares)
}