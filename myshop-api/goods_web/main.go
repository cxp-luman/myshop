package main

import (
	"fmt"
	"myshop-api/goods_web/global"
	"myshop-api/goods_web/initialize"
	"myshop-api/goods_web/utils"
	"myshop-api/goods_web/validators"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func main() {
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitRedisClient()
	initialize.InitSrvConn()
	err := initialize.InitTrans("zh")
	if err != nil {
		panic("init validator failed")
	}
	router := initialize.Routers()
	// todo: remove the validate mobile
	// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	// 	v.RegisterValidation("mobile", validators.ValidatorMobile)
	// 	_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
	// 		return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
	// 	}, func(ut ut.Translator, fe validator.FieldError) string {
	// 		t, _ := ut.T("mobile", fe.Field())
	// 		return t
	// 	})
	// } else {
	// 	zap.S().Error("Failed to convert validator type")
	// }
	slugger := zap.S()
	slugger.Debug("初始化gin路由")
	env := utils.GetEnvInfo("shop_env")
	port := 8080
	if env == "" || env == "dev" {
		port, err = utils.GetFreePort()
		if err != nil {
			panic("get free port failed")
		}
	}
	zap.S().Infow("init gin server", "port", port)
	err = router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		slugger.Fatal("初始化api服务失败:", err.Error())
	}

}
