package router

import (
	"github.com/gin-gonic/gin"
	"github.com/han-songyang/information-management-serve/infra/logger"
)

// InitUserRouter 初始化用户路由
func InitUserRouter() {

	userRouter := Router.Group("/user")
	{
		userRouter.GET("/info", func(ctx *gin.Context) {
			logger.InfoLog("userRouter", "info", "info")
			ctx.JSON(200, gin.H{
				"message": "user info",
			})
		})
	}
}
