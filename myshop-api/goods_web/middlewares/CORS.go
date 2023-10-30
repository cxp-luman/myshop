package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		ctx.Header("Access-Control-Allow-Origin", "*") // 表示允许所有源（即任何域名）的请求访问资源
		ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token") // 表示允许的请求头字段
		ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT") // 表示允许的请求方法
		ctx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type") //表示允许的请求方法
		ctx.Header("Access-Control-Allow-Credentials", "true") // 表示是否允许发送 Cookie

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
	}
}