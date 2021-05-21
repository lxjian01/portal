package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers/alarm"
)

func AlarmRoutes(route *gin.Engine) {
	user := route.Group("/api/portal/alarm")
	{
		// alarm group
		user.POST("/alarm_group", alarm.AddAlarmGroup)
		user.PUT("/alarm_group/:id", alarm.UpdateAlarmGroup)
		user.DELETE("/alarm_group/:id", alarm.DeleteAlarmGroup)
		user.GET("/alarm_group/list", alarm.GetAlarmGroupList)
		user.GET("/alarm_group/page", alarm.GetAlarmGroupPage)
	}
}