package middleware

import (
	"time"

	"github.com/fishjar/gin-boilerplate/logger"
	"github.com/fishjar/gin-boilerplate/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// LoggerToFile 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()                    // 开始时间
		reqInfo := service.GetReqInfo(c)           // 请求信息
		c.Next()                                   // 处理请求
		statusCode := c.Writer.Status()            // 返回状态码
		endTime := time.Now()                      // 结束时间
		latencyTime := endTime.Sub(startTime)      // 执行时间
		go logger.LogReq.WithFields(logrus.Fields{ // 日志记录
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    reqInfo.IP,
			"req_method":   reqInfo.Method,
			"req_uri":      reqInfo.URI,
		}).Info()
	}
}

// LoggerToMongo 日志记录到 MongoDB
func LoggerToMongo() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LoggerToES 日志记录到 ES
func LoggerToES() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// LoggerToMQ 日志记录到 MQ
func LoggerToMQ() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
