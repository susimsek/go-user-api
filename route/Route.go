package route

import (
	"github.com/gin-gonic/gin"
	"user-api/controller"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/user-api")
	{
		grp1.GET("user", controller.GetUsers)
		grp1.POST("user", controller.CreateUser)
		grp1.GET("user/:id", controller.GetUserByID)
		grp1.PUT("user/:id", controller.UpdateUser)
		grp1.DELETE("user/:id", controller.DeleteUser)
	}
	return r
}
