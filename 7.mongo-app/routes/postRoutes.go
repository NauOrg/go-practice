package routes

import (
	"mongo-project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(routes *gin.Engine) {
	postRoutesGroup := routes.Group("/post")
	{
		postRoutesGroup.GET("/:id", controllers.GetPostById)
		postRoutesGroup.GET("", controllers.GetPosts)
		postRoutesGroup.POST("", controllers.CreatePost)
		postRoutesGroup.DELETE("/:id", controllers.DeletePost)
	}
}
