package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("/user")
	{
		user.GET("/test", controllers.Test)
		//user.POST("/test", controllers.Test)
	}
}