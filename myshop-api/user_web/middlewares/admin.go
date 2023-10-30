package middlewares

import (
	"myshop-api/user_web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		claims, exists := ctx.Get("claims")
		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": "auth middllewares failed",
			})
			ctx.Abort()
		}
		usrInfo := claims.(*models.CustomClaims)
		if usrInfo.AuthorityId == 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "not has the admin auth",
			})
			ctx.Abort()
		}
		ctx.Next()

	}
} 