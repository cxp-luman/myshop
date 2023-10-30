package router

import (
	"myshop-api/user_web/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/getVerifyCode", api.GetVerifyCode)
		userRouter.POST("/sendTelCode", api.SmsSend)
	}
}