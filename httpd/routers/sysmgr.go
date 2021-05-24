package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers/sysmgr"
)

func SysmgrRoutes(route *gin.Engine) {
	user := route.Group("/api/portal/sysmgr")
	{
		// menu
		user.POST("/menu", sysmgr.AddMenu)
		user.PUT("/menu/:id", sysmgr.UpdateMenu)
		user.DELETE("/menu/:id", sysmgr.DeleteMenu)
		user.GET("/menu/list", sysmgr.GetMenuList)
		user.GET("/menu/parent_list", sysmgr.GetParentMenuList)
		user.GET("/menu/page", sysmgr.GetMenuPage)
		user.GET("/menu/:id", sysmgr.GetMenuDetail)

		// user
		user.POST("/user", sysmgr.AddUser)
		user.PUT("/user/:id", sysmgr.UpdateUser)
		user.DELETE("/user/:id", sysmgr.DeleteUser)
		user.GET("/user/page", sysmgr.GetUserPage)
		user.GET("/user/:id", sysmgr.GetUserDetail)

		// role
		user.POST("/role", sysmgr.AddRole)
		user.PUT("/role/:id", sysmgr.UpdateRole)
		user.DELETE("/role/:id", sysmgr.DeleteRole)
		user.GET("/role/list", sysmgr.GetRoleList)
		user.GET("/role/page", sysmgr.GetRolePage)
	}
}