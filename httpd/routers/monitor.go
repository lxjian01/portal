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
		user.GET("/cluster/list", monitor.GetMonitorClusterList)
		user.GET("/cluster/page", monitor.GetMonitorClusterPage)

		// monitor component
		user.POST("/component", monitor.AddMonitorComponent)
		user.PUT("/component/:id", monitor.UpdateMonitorComponent)
		user.DELETE("/component/:id", monitor.DeleteMonitorComponent)
		user.GET("/component/list", monitor.GetMonitorComponentList)
		user.GET("/component/page", monitor.GetMonitorComponentPage)

		// monitor target
		user.POST("/target", monitor.AddMonitorTarget)
		user.PUT("/target/:id", monitor.UpdateMonitorTarget)
		user.DELETE("/target/:id", monitor.DeleteMonitorTarget)
		user.GET("/target/page", monitor.GetMonitorTargetPage)
	}
}