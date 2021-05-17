package httpd

import (
	"github.com/gin-gonic/gin"
	"net"
	"portal/config"
	"portal/global/log"
	"portal/httpd/middlewares"
	"portal/httpd/routers"
	"strconv"
)

func StartHttpdServer(c *config.HttpdConfig) {
	router := gin.Default()
	// 添加自定义的 logger 间件
	router.Use(middlewares.Logger(), gin.Recovery())
	router.Use(middlewares.Auth(), gin.Recovery())
	// 添加路由
	routers.UserRoutes(router)      //Added all user routers
	routers.SysmgrRoutes(router)      //Added system mgr routers
	// 拼接host
	Host := c.Host
	Port := strconv.Itoa(c.Port)
	addr := net.JoinHostPort(Host, Port)
	log.Info("Start HTTP server at", addr)
	err1 := router.Run(addr)
	if err1 != nil{
		log.Error("Start server error by",err1)
	}
	log.Info("Start server ok")
}
