package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/han-songyang/information-management-serve/infra/conf"
	"io"
	"os"
)

var Router *gin.Engine

// InitRouter 初始化路由
func InitRouter() {

	f, _ := os.OpenFile("./log/gin.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	Router = gin.New()
	Router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - %s \"%s %s %s %d %s \"%s\" %s\"\n",
			param.TimeStamp.Format("2006-01-02 03:04:05"),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	InitUserRouter()

	fmt.Println("router init success")

	port := conf.Config.Section("server").Key("port").String()
	err := Router.Run(port)
	if err != nil {
		return
	}
}
