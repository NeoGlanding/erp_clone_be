package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/automa8e_clone/types"
	"github.com/gin-gonic/gin"
)

func Party(r *gin.Engine) {
	party := r.Group("/parties")

	party.GET("/", middlewares.TokenAuthenticationMiddleware, middlewares.PaginationMiddleware, middlewares.QueryMiddleware, controllers.GetParties, middlewares.ResponseMiddlewares)
	party.GET("/:id", middlewares.TokenAuthenticationMiddleware, controllers.GetParty, middlewares.ResponseMiddlewares)
	party.POST("/", middlewares.TokenAuthenticationMiddleware, controllers.PostParty, middlewares.ResponseMiddlewares)
	party.PUT("/:id", middlewares.TokenAuthenticationMiddleware, middlewares.InterceptParam("id", "party-id"),middlewares.PartyAuthorizationRole([]string{types.PERMISSION_OWNER}),controllers.UpdateParty, middlewares.ResponseMiddlewares)

	party.POST("/action", middlewares.TokenAuthenticationMiddleware, controllers.PartyAction,middlewares.ResponseMiddlewares)
}