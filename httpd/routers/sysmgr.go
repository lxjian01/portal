package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers/sysmgr"
)

func SysmgrRoutes(route *gin.Engine) {
	user := route.Group("/api/cmdb/sysmgr")
	{
		// system menu
		user.POST("/menu", sysmgr.AddMenu)
		user.PUT("/menu/:uid", sysmgr.UpdateMenu)
		user.DELETE("/menu/:uid", sysmgr.DeleteMenu)
		user.GET("/menu/list", sysmgr.GetMenuList)
		user.GET("/menu/page", sysmgr.GetMenuPage)
		user.GET("/menu/detail/:id", sysmgr.GetMenuDetail)
	}
}