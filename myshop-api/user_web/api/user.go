package api

import (
	"context"
	"fmt"
	"myshop-api/user_web/global"
	"myshop-api/user_web/middlewares"
	"myshop-api/user_web/models"
	"myshop-api/user_web/proto"
	"myshop-api/user_web/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetUserList(ctx *gin.Context) {
	s := zap.S()
	claims, exists := ctx.Get("claims")
	if !exists {
		s.Error("get claims failed")
		panic("get claims failed")
	}
	userInfo := claims.(*models.CustomClaims)
	s.Infof("user %d use", userInfo.ID)
	// 调用rpc服务
	pageNumber, _ := strconv.Atoi(ctx.DefaultQuery("pN", "0"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pSize", "10"))
	userListResp, err := global.Srvclient.GetUserList(context.Background(), &proto.PageInfo{
		PN:    uint32(pageNumber),
		PSize: uint32(pageSize),
	})
	if err != nil {
		s.Errorw("[GetUserList]", "err", err.Error())
		utils.CodeTransFromGrpcToHttp(err, ctx)
		return
	}
	usersInfo := make([]interface{}, 0)
	for _, user := range userListResp.Data {
		user := global.UserResponse{
			Id:       user.Id,
			Mobile:   user.Mobile,
			NickName: user.NickName,
			Birthday: global.JsonTime(time.Unix(int64(user.BirthDay), 0)),
			Gender:   user.Gender,
		}
		usersInfo = append(usersInfo, user)
	}
	resp := gin.H{
		"total":    userListResp.Total,
		"userList": usersInfo,
	}
	ctx.JSON(http.StatusOK, resp)
	s.Infow("[getUserList]", "mes", "get userInfo success!")
}

func PassWord(ctx *gin.Context) {
	passWordLoginReq := global.PassWordLoginReq{}
	if err := ctx.ShouldBind(&passWordLoginReq); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}

	if checkVerifyCode := strore.Verify(passWordLoginReq.VerifyCodeId, passWordLoginReq.VerifyCode, true); !checkVerifyCode {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "Verification code error",
		})
		return
	}
	connect, err := grpc.Dial(fmt.Sprintf("%s:%s", global.UserWebInfo.UserServerInfo.Host, global.UserWebInfo.UserServerInfo.Port), grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Dial user srv failed",
		})
		return
	}
	userSrvClient := proto.NewUserClient(connect)
	userResp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passWordLoginReq.Mobile,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			switch s.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "Network Error",
				})
			}
			return
		}
	} else {
		if check, err := userSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			Password: passWordLoginReq.PassWord, EncryptedPassword: userResp.PassWord,
		}); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"meg": "Network Error",
			})
			return
		} else {
			if check.Success {
				fmt.Println(userResp.Role)
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(userResp.Id),
					NickName:    userResp.NickName,
					AuthorityId: uint(userResp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: int64(time.Now().Unix()),
						ExpiresAt: time.Now().Unix() + 60*60*24*15,
						Issuer:    "cxp",
					},
				}
				token, err := j.CreateToken(claims)
				fmt.Println(err)
				if err != nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{
						"msg": "create token failed",
					})
					return
				}

				ctx.JSON(http.StatusOK, gin.H{
					"id":         userResp.Id,
					"nick_name":  userResp.NickName,
					"token":      token,
					"expires_at": time.Now().Unix() + 60*60*24*15,
				})
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"msg": "password is nor corret!",
				})
			}
		}
	}
}

func RegisterUser(ctx *gin.Context) {
	registerReq := global.RegisterReq{}
	if err := ctx.ShouldBind(&registerReq); err != nil {
		utils.HandleValidatorError(ctx, err)
		return
	}
	val, err := global.Rdb.Get(context.Background(),registerReq.Mobile).Result()
	if err != nil {
		panic(err)
	}
	if err == redis.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "the tel verify code expired",
		})
		return
	}
	if val != registerReq.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":"验证码错误",
		})
		return
	}
	connect, err := grpc.Dial(fmt.Sprintf("%s:%s", global.UserWebInfo.UserServerInfo.Host, global.UserWebInfo.UserServerInfo.Port), grpc.WithInsecure())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "Dial user srv failed",
		})
		return
	}
	userSrvClient := proto.NewUserClient(connect)
	registerResp, err := userSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerReq.Mobile,
		PassWord: registerReq.PassWord,
		Mobile:   registerReq.Mobile,
	})
	if err != nil {
		zap.S().Errorw("reegister failed", "mobile", registerReq.Mobile)
		utils.CodeTransFromGrpcToHttp(err, ctx)
		return
	}
	j := middlewares.NewJWT()
	if token, err := j.CreateToken(models.CustomClaims{
		ID: uint(registerResp.Id),
		NickName: registerResp.NickName,
		AuthorityId: uint(registerResp.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 3600*60*15,
			Issuer: "cxp",
		},
	}); err != nil {
		zap.S().Errorw("create token failed", "mobile", registerReq.Mobile)
		panic("create token failed")
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id": registerResp.Id,
			"nick_name": registerResp.NickName,
			"token": token,
			"expired_at": (time.Now().Unix() + 60*60*24*30)*1000,
		})
	}
	
}