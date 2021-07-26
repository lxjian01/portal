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

		// monitor prometheus
		user.POST("/prometheus", monitor.AddMonitorPrometheus)
		user.PUT("/prometheus/:id", monitor.UpdateMonitorPrometheus)
		user.DELETE("/prometheus/:id", monitor.DeleteMonitorPrometheus)
		user.GET("/prometheus/list", monitor.GetMonitorPrometheusList)
		user.GET("/prometheus/page", monitor.GetMonitorPrometheusPage)

		// monitor resource
		user.POST("/resource", monitor.AddMonitorResource)
		user.PUT("/resource/:id", monitor.UpdateMonitorResource)
		user.DELETE("/resource/:id", monitor.DeleteMonitorResource)
		user.GET("/resource/list", monitor.GetMonitorResourceList)
		user.GET("/resource/page", monitor.GetMonitorResourcePage)

		// monitor target
		user.POST("/target", monitor.AddMonitorTarget)
		user.PUT("/target/:id", monitor.UpdateMonitorTarget)
		user.DELETE("/target/:id", monitor.DeleteMonitorTarget)
		user.GET("/target/page", monitor.GetMonitorTargetPage)
	}
}