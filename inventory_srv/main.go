package main

import (
	"flag"
	"fmt"
	"myshop/inventory_srv/global"
	"myshop/inventory_srv/handler"
	"myshop/inventory_srv/initialize"
	"myshop/inventory_srv/proto"
	"myshop/inventory_srv/utils"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	IP := flag.String("IP", "0.0.0.0", "please input ip")
	PORT := flag.Int("PORT", 0, "please input port")
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	var err error
	if *PORT == 0 {
		if *PORT, err = utils.GetFreePort(); err != nil {
			zap.S().Fatal("get random port failed")
		}
		zap.S().Infof("port: %d", *PORT)
	}
	*PORT = 50051
	fmt.Println(*PORT)
	server := grpc.NewServer()
	proto.RegisterInventoryServer(server, &handler.InventoryServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *PORT))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.InvSrvInfo.ConsulInfo.Host, global.InvSrvInfo.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("192.168.1.100:%d", *PORT),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	namespace := uuid.New()
	name := []byte("user_srv")
	userSrvUuid := uuid.NewSHA1(namespace, name)
	serviceId := userSrvUuid.String()
	registration := new(api.AgentServiceRegistration)
	registration.Address = "192.168.1.100"
	registration.Port = *PORT
	registration.Name = global.InvSrvInfo.Name
	registration.ID = serviceId
	registration.Tags = []string{global.InvSrvInfo.Name}
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serviceId); err != nil{
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
