package initialize

import (
	"fmt"
	"myshop-api/user_web/global"
	"myshop-api/user_web/proto"

	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
)
func InitSrvConn() {
	consulInfo := global.UserWebInfo.ConsulInfo
	consul_lb_cfg := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.UserWebInfo.UserServerInfo.Name)
	conn, err := grpc.Dial(
        consul_lb_cfg,
        grpc.WithInsecure(),
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
    )
    if err != nil {
		zap.S().Fatal("load balance get rpc connect failed")
		panic(err)
    }
	// cfg := api.DefaultConfig()
	// cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	// conculClient, err := api.NewClient(cfg)
	// s := zap.S()
	// if err != nil {
	// 	s.Fatal("connect consul failed")
	// 	panic(err)
	// }
	// userSrvInfo, err := conculClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.UserWebInfo.UserServerInfo.Name))
	// if err != nil {
	// 	zap.S().Fatal("get srv info failed")
	// 	panic(err)
	// }
	// consulHost := ""
	// consulPort := 0
	// for _, value := range userSrvInfo {
	// 	consulHost = value.Address
	// 	consulPort = value.Port
	// }
	// // todo: modify the deprecated grpc.WithInsecure
	// conn, err := grpc.Dial(fmt.Sprintf("%s:%d", consulHost, consulPort), grpc.WithInsecure())
	// if err != nil {
	// 	s.Fatalw("get grpc client channel failed", "err", err)
	// }
	client := proto.NewUserClient(conn)
	global.Srvclient = client
}

func InitSrvConnOld() {
	consulInfo := global.UserWebInfo.ConsulInfo
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	conculClient, err := api.NewClient(cfg)
	s := zap.S()
	if err != nil {
		s.Fatal("connect consul failed")
		panic(err)
	}
	userSrvInfo, err := conculClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.UserWebInfo.UserServerInfo.Name))
	if err != nil {
		zap.S().Fatal("get srv info failed")
		panic(err)
	}
	consulHost := ""
	consulPort := 0
	for _, value := range userSrvInfo {
		consulHost = value.Address
		consulPort = value.Port
	}
	// todo: modify the deprecated grpc.WithInsecure
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", consulHost, consulPort), grpc.WithInsecure())
	if err != nil {
		s.Fatalw("get grpc client channel failed", "err", err)
	}
	client := proto.NewUserClient(conn)
	global.Srvclient = client
}