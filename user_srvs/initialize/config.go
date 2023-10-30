package initialize

import (
	"encoding/json"
	"myshop/user_srvs/global"
	"time"

	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func InitConfig() {
	v := viper.New()
	v.SetConfigFile("config-debug.yaml")
	if err := v.ReadInConfig(); err != nil {
		zap.S().Fatal("init config failed!")
		panic(err)
	}
	if err := v.Unmarshal(&global.NacosInfo);err != nil {
		zap.S().Fatal("unmashal usersrvinfo failed!")
		panic(err)
	}
	fmt.Println(global.NacosInfo)
	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
			zap.S().Info("the config file changed")
			err := v.ReadInConfig()
			if err != nil {
				panic(err)
			}
			err = v.Unmarshal(global.UserSrvInfo)
			if err != nil {
				panic(err)
			}
		})
		time.Sleep(time.Second * 3600)
	}()
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "1d17067d-0edc-4dc2-ab5f-afca156200ce", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      "http://10.211.55.10",
			ContextPath: "/nacos",
			Port:        8848,
			Scheme:      "http",
		},
	}
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user_srv",
		Group:  "user"})
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(content), &global.UserSrvInfo)
}
