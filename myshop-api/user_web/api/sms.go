package api

import (
	"context"
	"fmt"
	"math/rand"
	"myshop-api/user_web/global"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func GenerateSmsCode(witdh int) string {
	//生成width长度的短信验证码

	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < witdh; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SmsSend(ctx *gin.Context) {
	sendSmsReq := global.SendSmsReq{}
	if err := ctx.ShouldBind(&sendSmsReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "please write corrrect filed",
		})
		return
	}
	s := GenerateSmsCode(6)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": s,
	})
	rdb := global.Rdb
	s2, err2 := rdb.Get(context.Background(), "test-key").Result()
	fmt.Println(err2)
	if err2 != nil {
		fmt.Println(s2)
	}
	fmt.Println(rdb)
	err := rdb.Set(context.Background(), sendSmsReq.Mobile, s,60*time.Second).Err()
	if err != nil {
		panic(err)
	}
}