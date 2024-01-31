package routes

import (
	"github.com/automa8e_clone/controllers"
	"github.com/automa8e_clone/middlewares"
	"github.com/gin-gonic/gin"
)

func Files(r *gin.Engine) {
	files := r.Group("/files")

	files.GET("/:id", middlewares.TokenAuthenticationMiddleware, controllers.GetFile, middlewares.ResponseMiddlewares)
	files.POST("/", middlewares.TokenAuthenticationMiddleware, controllers.PostFile,middlewares.ResponseMiddlewares)

}