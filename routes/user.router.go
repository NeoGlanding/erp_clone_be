package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Users(r *gin.Engine) {
	users := r.Group("/users")

	users.POST("/onboard", middlewares.TokenAuthenticationMiddleware, middlewares.UnonboardedAuthorization, middlewares.BodyCountryIdExistMiddleware, middlewares.InterceptFileIdBody("profile_picture_file_id"),middlewares.FileIdExist, controllers.OnboardUser,middlewares.ResponseMiddlewares)
	users.PUT("/", middlewares.TokenAuthenticationMiddleware, controllers.UpdateCredential, middlewares.ResponseMiddlewares)
}