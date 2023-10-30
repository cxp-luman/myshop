package initialize

import (
	"myshop-api/goods_web/middlewares"
	"myshop-api/goods_web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	apiRouter := Router.Group("/v1")
	router.InitUserRouter(apiRouter)
	return Router
}
