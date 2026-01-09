package routes

import (
	"mongo-project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(routes *gin.Engine) {
	userRoutesGroup := routes.Group("/user")
	{
		userRoutesGroup.GET("", controllers.GetUsers)
		userRoutesGroup.GET("/:id", controllers.GetUserById)
		userRoutesGroup.POST("", controllers.CreateUser)
		userRoutesGroup.DELETE("/:id", controllers.DeleteUser)
	}
}
