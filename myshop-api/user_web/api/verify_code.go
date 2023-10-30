package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var strore = base64Captcha.DefaultMemStore

func GetVerifyCode(ctx *gin.Context) {
	driver := base64Captcha.NewDriverDigit(80, 240, 6, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, strore)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("生成验证码错误,: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "generate chaptcha failed",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
		"chaptcha": b64s,
	})
}