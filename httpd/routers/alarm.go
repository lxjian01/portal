package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers/alarm"
)

func AlarmRoutes(route *gin.Engine) {
	user := route.Group("/api/portal/alarm")
	{
		// alarm group
		user.POST("/group", alarm.AddAlarmGroup)
		user.PUT("/group/:id", alarm.UpdateAlarmGroup)
		user.DELETE("/group/:id", alarm.DeleteAlarmGroup)
		user.GET("/group/page", alarm.GetAlarmGroupPage)
		user.GET("/group/list", alarm.GetAlarmGroupList)
		user.GET("/group/:id", alarm.GetAlarmGroupDetail)

		// alarm user
		user.POST("/user", alarm.AddAlarmUser)
		user.PUT("/user/:id", alarm.UpdateAlarmUser)
		user.DELETE("/user/:id", alarm.DeleteAlarmUser)
		user.GET("/user/list", alarm.GetAlarmUserList)
		user.GET("/user/page", alarm.GetAlarmUserPage)
	}
}