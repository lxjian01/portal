package routers

import (
	"github.com/gin-gonic/gin"
	"portal/httpd/controllers/monitor"
)

func MonitorRoutes(route *gin.Engine) {
	user := route.Group("/api/portal/monitor")
	{
		// monitor cluster
		user.POST("/cluster", monitor.AddMonitorCluster)
		user.PUT("/cluster/:id", monitor.UpdateMonitorCluster)
		user.DELETE("/cluster/:id", monitor.DeleteMonitorCluster)
		user.GET("/cluster/page", monitor.GetMonitorClusterPage)
		user.GET("/cluster/list", monitor.GetMonitorClusterList)
	}
}