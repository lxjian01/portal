package web

import (
	"github.com/gin-gonic/gin"
	"net"
	"portal/global/config"
	"portal/global/log"
	"portal/web/middlewares"
	"portal/web/routers"
	"strconv"
)

func StartServer(c *config.HttpdConfig) {
	router := gin.Default()
	// 添加自定义的 logger 间件
	router.Use(middlewares.Logger(), gin.Recovery())
	router.Use(middlewares.Auth(), gin.Recovery())
	// 添加路由
	routers.SysmgrRoutes(router)      //Added system mgr routers
	routers.MonitorRoutes(router)      //Added monitor routers
	routers.AlarmRoutes(router)      //Added alarm routers
	// 拼接host
	Host := c.Host
	Port := strconv.Itoa(c.Port)
	addr := net.JoinHostPort(Host, Port)
	log.Infof("Start HTTP server at %s", addr)
	err := router.Run(addr)
	if err != nil{
		log.Errorf("Start server error by %v",err)
	}
	log.Info("Start server ok")
}
