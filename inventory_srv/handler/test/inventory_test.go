package test_test

import (
	"context"
	"myshop/inventory_srv/proto"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var invClient proto.InventoryClient
var conn *grpc.ClientConn

func Init(){
	var err error
	conn, err = grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	invClient = proto.NewInventoryClient(conn)
}

func TestSetInv(T *testing.T) {
	Init()
	_, err := invClient.SetInv(context.Background(), &proto.GoodsInvInfo{GoodsId: 421, Num: 100})
	if err != nil {
		panic(err)
	}
}