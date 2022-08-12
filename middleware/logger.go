package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/gin-blog"
	//linkName := "letestlog.log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755) // 创建一个log日志文件
	if err != nil {
		fmt.Println("err:", err)
	}
	//由于logrus本身并不提供写文件, 并且按照日期自动分割, 删除过期日志文件的功能.
	//一般情况下大家都是使用github.com/rifflock/lfshook配合github.com/lestrrat-go/file-rotatelogs来实现相关的功能
	logger := logrus.New()             //实例化
	logger.Out = scr                   //设置log文件输出 输出到文件，不在控制台输出
	logger.SetLevel(logrus.DebugLevel) //设置日志级别

	writer, _ := retalog.New( //自动切割日志文件
		filePath+"%Y%m%d.log",                    //文件名
		retalog.WithMaxAge(7*24*time.Hour),       //保留时间
		retalog.WithRotationTime(24*time.Minute), //多长时间分割一次
		//retaLog.WithLinkName(linkName),		//软链接
	)
	lfsHook := lfshook.WriterMap{ //WriterMap是将日志级别映射到 io.Writer 的映射
		logrus.DebugLevel: writer,
		logrus.ErrorLevel: writer,
		logrus.FatalLevel: writer,
		logrus.PanicLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.InfoLevel:  writer,
		logrus.TraceLevel: writer,
	}
	//
	Hook := lfshook.NewHook(lfsHook, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	//调用AddHook时, 将Hook加入到LevelHooks map中
	//程序打印log, 会最终执行到Entry.log()
	//Entry.log()会调用fireHooks()
	//fireHooks又会调用LevelHooks Fire()函数, 该函数会遍历所有的Hook, 从而执行相应的Hook
	logger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		//gin里面是洋葱模型中间件，先执行c.Next上面的内容，在执行下一个中间件，下一个中间件也是执行c.Next上面的内容，都执行完后返回到第一个中间件c.Next后面的内容
		c.Next()
		//一个时间段
		stopTime := time.Since(startTime).Milliseconds()
		//开销时间为一个整数  向上取整 stopTime的纳秒  格式化输出字符串
		spendTime := fmt.Sprintf("%d ms", stopTime)
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
			entry.Info()
		}
	}
}
