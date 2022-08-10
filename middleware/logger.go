package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()
	return func(c *gin.Context) {
		startTime := time.Now()
		//gin里面是洋葱模型中间件，先执行c.Next上面的内容，在执行下一个中间件，下一个中间件也是执行c.Next上面的内容，都执行完后返回到第一个中间件c.Next后面的内容
		c.Next()
		//一个时间段
		stopTime := time.Since(startTime)
		//开销时间为一个整数  向上取整 stopTime的纳秒  格式化输出字符串
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()))/1000000.0))
		//请求客户端hostname
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"stats":     statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info(time.Now())
		}
	}
}
