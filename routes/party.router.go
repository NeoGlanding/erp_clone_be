package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Party(r *gin.Engine) {
	party := r.Group("/parties")

	party.GET("/", middlewares.TokenAuthenticationMiddleware, middlewares.PaginationMiddleware, middlewares.QueryMiddleware, controllers.GetParties, middlewares.ResponseMiddlewares)
	party.POST("/", middlewares.TokenAuthenticationMiddleware, controllers.PostParty, middlewares.ResponseMiddlewares)
	party.PUT("/:id", middlewares.TokenAuthenticationMiddleware, controllers.UpdateParty, middlewares.ResponseMiddlewares)
}