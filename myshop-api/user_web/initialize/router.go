package initialize

import (
	"myshop-api/user_web/middlewares"
	"myshop-api/user_web/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	apiRouter := Router.Group("/v1")
	router.InitBaseRouter(apiRouter)
	router.InitUserRouter(apiRouter)
	return Router
}
