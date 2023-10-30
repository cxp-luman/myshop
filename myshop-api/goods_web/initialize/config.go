package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"myshop-api/goods_web/global"
	"time"
)

func getEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	envInfo := getEnvInfo("test")
	var configFilePath string
	if envInfo == "test" {
		configFilePath = fmt.Sprintf("../%s", "config-test.yaml")
	} else {
		configFilePath = fmt.Sprintf("%s", "config-test.yaml")
	}
	v := viper.New()
	v.SetConfigFile(configFilePath)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Info(err)
		fmt.Println(err)
		panic("get the config file failed")
	}
	err := v.Unmarshal(global.NacosInfo)
	if err != nil {
		panic(err)
	}
	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
			zap.S().Info("the config file changed")
			err = v.ReadInConfig()
			if err != nil {
				panic(err)
			}
			err = v.Unmarshal(global.NacosInfo)
			if err != nil {
				panic(err)
			}
		})
		time.Sleep(time.Second * 3600)
	}()
}
