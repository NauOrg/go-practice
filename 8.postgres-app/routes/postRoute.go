package routes

import (
	"postgres-project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPost(r *gin.Engine) {
	post := r.Group("/post")
	{
		post.GET("", controllers.GetPostAll)
		post.GET("/:id", controllers.GetPostById)
		post.POST("", controllers.CreatePost)
		post.DELETE("/:id", controllers.DeletePost)
	}
}

// {
//     "title": "test title3",
//     "user_id":7
// }
