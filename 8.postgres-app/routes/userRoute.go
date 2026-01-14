package routes

import (
	"postgres-project/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUser(r *gin.Engine) {
	user := r.Group("/user")
	{
		user.GET("", controllers.GetUserAll)
		user.GET("/:id", controllers.GetUserById)
		user.POST("", controllers.CreateUser)
		user.PUT("/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}
}
