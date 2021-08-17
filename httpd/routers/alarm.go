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

		// alarm recording rule
		user.POST("/recording/rule", alarm.AddRecordingRule)
		user.PUT("/recording/rule/:id", alarm.UpdateRecordingRule)
		user.DELETE("/recording/rule/:id", alarm.DeleteRecordingRule)
		user.GET("/recording/rule/page", alarm.GetRecordingRulePage)

		// alarm alerting metric
		user.POST("/alerting/metric", alarm.AddAlertingMetric)
		user.PUT("/alerting/metric/:id", alarm.UpdateAlertingMetric)
		user.DELETE("/alerting/metric/:id", alarm.DeleteAlertingMetric)
		user.GET("/alerting/metric/list", alarm.GetAlertingMetricList)
		user.GET("/alerting/metric/page", alarm.GetAlertingMetricPage)

		// alarm alerting rule
		user.GET("/alerting/rule/page", alarm.GetAlertingRulePage)

		// alarm page
		user.POST("/notice", alarm.AddAlarmNotice)
		user.GET("/notice/page", alarm.GetAlarmNoticePage)
	}
}