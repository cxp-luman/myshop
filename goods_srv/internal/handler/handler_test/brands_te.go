package main

import (
	"context"
	"fmt"
	"myshop/goods_srv/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func ABrandList() {
	cc, err := grpc.Dial("localhost:51790", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}
	defer cc.Close()
	gc := proto.NewGoodsClient(cc)
	blr, err := gc.BrandList(context.Background(), &proto.BrandFilterRequest{Pages: 1, PagePerNums: 10})
	fmt.Println(blr)
	if err != nil {
		panic(err)
	}

	fmt.Println(blr)

}

func ACreateBrand() {
	fmt.Println("dasdsa")
	cc, err := grpc.Dial("localhost:51790", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)

	}
	defer cc.Close()
	gc := proto.NewGoodsClient(cc)
	bir, err := gc.CreateBrand(context.Background(), &proto.BrandRequest{Id: 11111, Name: "TEST_ONE", Logo: "HTTTP:FSDSFSDFDSFSDFDFSDFDSFDF"})
	if err != nil {
		panic(err)
	}
	fmt.Println(bir)

}

