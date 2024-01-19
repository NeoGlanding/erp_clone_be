package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	auth := r.Group("/auth")

	auth.POST("/login", controllers.Login)
}