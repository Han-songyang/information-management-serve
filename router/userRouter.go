package router

import (
	"github.com/gin-gonic/gin"
)

// InitUserRouter 初始化用户路由
func InitUserRouter() {

	userRouter := Router.Group("/user")
	{
		userRouter.GET("/info", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "user info",
			})
		})
	}
}
