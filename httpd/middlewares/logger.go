package middlewares

import (
	"github.com/gin-gonic/gin"
	"portal/global/log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(context *gin.Context) {
		host := context.Request.Host
		url := context.Request.URL
		method := context.Request.Method
		startTime := time.Now().Format("2006-01-02 15:04:05")
		log.Infof("%s %s %s%s \n",startTime, method, host, url)
		context.Next()
		endTime := time.Now().Format("2006-01-02 15:04:05")
		log.Infof("%s %s %s%s response status is %d \n",endTime, method, host, url,context.Writer.Status())
	}
}
