package utils

import (
	"fmt"
	"myshop-api/goods_web/global"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CodeTransFromGrpcToHttp(err error, ctx *gin.Context) {
	if err != nil {
		e, _ := status.FromError(err)
		switch e.Code() {
		case codes.NotFound:
			ctx.JSON(http.StatusNotFound, gin.H{
				"msg": e.Message(),
			})
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"meg": e.Message(),
			})
		case codes.Internal:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": e.Message(),
			})
		case codes.Unavailable:
			ctx.JSON(http.StatusBadGateway, gin.H{
				"msg": e.Message(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg": e.Message(),
			})
		}
	}
}

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		fmt.Println("failed!")
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}

func GetFreePort() (int, error) {
    addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
    if err != nil {
        return 0, err
    }

    l, err := net.ListenTCP("tcp", addr)
    if err != nil {
        return 0, err
    }
    defer l.Close()
    return l.Addr().(*net.TCPAddr).Port,  nil
}
func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}