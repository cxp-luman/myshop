package initialize

import (
	"fmt"
	"myshop-api/goods_web/global"
	"myshop-api/goods_web/proto"

	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func InitSrvConn() {
	consulInfo := global.GoodsWebInfo.ConsulInfo
	consul_lb_cfg := fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.GoodsWebInfo.GoodsServerInfo.Name)
	// 创建不安全的凭证（insecure credentials）
	creds := insecure.NewCredentials()
	// 使用WithTransportCredentials将凭证应用到连接选项中
	connOpts := grpc.WithTransportCredentials(creds)

	conn, err := grpc.Dial(
        consul_lb_cfg,
		connOpts,
        grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
    )
    if err != nil {
		zap.S().Fatal("load balance get rpc connect failed")
		panic(err)
    }
	client := proto.NewGoodsClient(conn)
	global.Srvclient = client
}

func InitSrvConnOld() {
	consulInfo := global.GoodsWebInfo.ConsulInfo
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)
	conculClient, err := api.NewClient(cfg)
	s := zap.S()
	if err != nil {
		s.Fatal("connect consul failed")
		panic(err)
	}
	userSrvInfo, err := conculClient.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.GoodsWebInfo.GoodsServerInfo.Name))
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
	client := proto.NewGoodsClient(conn)
	global.Srvclient = client
}