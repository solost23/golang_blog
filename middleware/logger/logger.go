package logger

import (
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	logger := initLogger()
	return func(c echo.Context) error {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		next(c)
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request().Method
		// 请求路由
		reqUri := c.Request().RequestURI
		// 状态码
		statusCode := c.Response().Status
		// 请求ip
		clientIP := c.RealIP()
		// 日志格式
		logger.Infof("| %3d | %13v | %15s | %s| %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
		return nil
	}
}

// 初始化Logger
func initLogger() *logrus.Logger {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/log/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		log.Println(err.Error())
	}
	logFileName := now.Format("2006-01-02") + ".log"
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err.Error())
		}
	}
	// 写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err.Error())
	}
	// 实例化
	logger := logrus.New()
	// 设置输出
	logger.Out = src
	// 设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	// 设置日志格式
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	return logger
}
