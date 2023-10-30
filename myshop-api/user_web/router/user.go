package router

import (
	"myshop-api/user_web/api"
	"myshop-api/user_web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/getUserList", middlewares.JWTAuth(), middlewares.AuthAdmin(), api.GetUserList)
		userRouter.POST("/loginPw", api.PassWord)
		userRouter.POST("/register", api.RegisterUser)
	}
}
