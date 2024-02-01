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
	party.GET("/:id/users", middlewares.TokenAuthenticationMiddleware, middlewares.PaginationMiddleware,middlewares.QueryMiddleware,middlewares.InterceptParam("id", "party-id"),middlewares.PartyAuthorizationRole([]string{types.PERMISSION_ADMIN, types.PERMISSION_OWNER, types.PERMISSION_VIEWER}), controllers.GetPartyUsers,middlewares.ResponseMiddlewares)
	party.POST("/", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, middlewares.InterceptFileIdBody("file_id"),middlewares.FileIdExist, middlewares.BodyCountryIdExistMiddleware,controllers.PostParty, middlewares.ResponseMiddlewares)
	party.PUT("/:id", middlewares.TokenAuthenticationMiddleware, middlewares.OnboardedAuthorization, middlewares.InterceptParam("id", "party-id"),middlewares.PartyAuthorizationRole([]string{types.PERMISSION_OWNER}),middlewares.BodyCountryIdExistMiddleware,controllers.UpdateParty, middlewares.ResponseMiddlewares)

	party.POST("/action", middlewares.TokenAuthenticationMiddleware, middlewares.InterceptPartyIdFromBody, middlewares.PartyAuthorizationRole([]string{types.PERMISSION_OWNER, types.PERMISSION_ADMIN}),controllers.PartyAction,middlewares.ResponseMiddlewares)
}