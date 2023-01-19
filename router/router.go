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

	f, _ := os.Create("./log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	Router = gin.Default()

	InitUserRouter()

	fmt.Println("router init success")

	port := conf.Config.Section("server").Key("port").String()
	err := Router.Run(port)
	if err != nil {
		return
	}
}
