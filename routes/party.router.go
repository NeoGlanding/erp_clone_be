package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Party(r *gin.Engine) {
	party := r.Group("/parties")

	party.GET("/", middlewares.TokenAuthenticationMiddleware, controllers.GetParty)
}