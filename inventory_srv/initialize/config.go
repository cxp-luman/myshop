package initialize

import (
	"encoding/json"
	"myshop/inventory_srv/global"
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
	nacosInfo := global.NacosInfo
	go func() {
		v.WatchConfig()
		v.OnConfigChange(func(in fsnotify.Event) {
			zap.S().Info("the config file changed")
			err := v.ReadInConfig()
			if err != nil {
				panic(err)
			}
			err = v.Unmarshal(global.InvSrvInfo)
			if err != nil {
				panic(err)
			}
		})
		time.Sleep(time.Second * 3600)
	}()
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         nacosInfo.NameSpaceId, //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}
	// At least one ServerConfig
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosInfo.Ipaddr,
			ContextPath: "/nacos",
			Port:        uint64(nacosInfo.Port),
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
		DataId: nacosInfo.DataId,
		Group:  nacosInfo.Group})
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(content), &global.InvSrvInfo)
}
