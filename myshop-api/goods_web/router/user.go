package router

import (

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	goodsRouter := router.Group("/goods")
	{
		goodsRouter.POST("/list", )

	}
}
