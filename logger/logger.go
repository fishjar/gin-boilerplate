/*
	日志记录器
*/

package logger

import (
	"io"
	"os"
	"path"
	"time"

	"github.com/fishjar/gin-boilerplate/config"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

// LogFile 日志文件
var LogFile *os.File

// LogGinFile 日志文件
var LogGinFile *os.File

// LogReqFile 日志文件
var LogReqFile *os.File

// LogGormFile 日志文件
var LogGormFile *os.File

// Log 日志记录器
var Log = logrus.New()

// LogReq 日志记录器
var LogReq = logrus.New()

// LogGorm 日志记录器
var LogGorm = logrus.New()

func init() {
	// 获取日志路径
	rootDir, _ := os.Getwd()
	logDir := path.Join(rootDir, "tmp/log")

	// 创建日志路径
	err := os.MkdirAll(logDir, 0755)
	if err != nil {
		panic("创建日志目录失败")
	}

	// 创建日志文件
	LogFile, err := rotatelogs.New(
		path.Join(logDir, "log.%Y%m%d.log"),
		rotatelogs.WithLinkName(path.Join(logDir, "log.log")),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("创建日志文件失败")
	}

	// 创建GIN日志文件
	LogGinFile, err := rotatelogs.New(
		path.Join(logDir, "gin.%Y%m%d.log"),
		rotatelogs.WithLinkName(path.Join(logDir, "gin.log")),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("创建GIN日志文件失败")
	}

	// 创建REQ日志文件
	LogReqFile, err := rotatelogs.New(
		path.Join(logDir, "req.%Y%m%d.log"),
		rotatelogs.WithLinkName(path.Join(logDir, "req.log")),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("创建REQ日志文件失败")
	}

	// 创建GORM日志文件
	LogGormFile, err := rotatelogs.New(
		path.Join(logDir, "gorm.%Y%m%d.log"),
		rotatelogs.WithLinkName(path.Join(logDir, "gorm.log")),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic("创建GORM日志文件失败")
	}

	// 配置日志记录器文件
	Log.Out = LogFile
	LogReq.Out = LogReqFile
	LogGorm.Out = LogGormFile

	// 配置GIN日志文件
	if config.Config.APPEnv == "development" {
		gin.DefaultWriter = io.MultiWriter(LogGinFile, os.Stdout)
	} else {
		gin.DefaultWriter = io.MultiWriter(LogGinFile)
	}
	// TODO: 定义路由日志的格式 gin.DebugPrintRouteFunc

}
