package handler

import (
	"context"
	"myshop/inventory_srv/global"
	"myshop/inventory_srv/model"
	"myshop/inventory_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func (s *InventoryServer) SetInv(ctx context.Context, req *proto.GoodsInvInfo) (*emptypb.Empty, error) {
	db := global.DB
	inv := model.Inventory{}
	// 将查询出的id值设置到inv结构体中 save时有id则update无则insert
	global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv)
	inv.Goods = req.GoodsId
	inv.Stocks = req.Num
	result := db.Save(&inv)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, result.Error.Error())
	}
	return &emptypb.Empty{}, nil
}

func (s *InventoryServer) InvDetail(ctx context.Context, req *proto.GoodsInvInfo) (*proto.GoodsInvInfo, error) {
	var inv model.Inventory
	if result := global.DB.Where(&model.Inventory{Goods: req.GoodsId}).First(&inv); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "inventory info is not exist")
	}
	return &proto.GoodsInvInfo{
		GoodsId: inv.Goods,
		Num:     inv.Stocks,
	}, nil
}

func (s *InventoryServer) Sell(ctx context.Context, req *proto.SellInfo) (*emptypb.Empty, error) {
	tx := db.Begin()
	for _, goodInfo := range req.GoodsInfo {
		inv := model.Inventory{}
		if res := global.DB.Where(&model.Inventory{Goods: goodInfo.GoodsId}).First(&inv); res.RowsAffected == 0 {
			return nil, status.Errorf(codes.InvalidArgument, "inventory info not exist")
		}
		if inv.Stocks < goodInfo.Num {
			return nil, status.Errorf(codes.ResourceExhausted, "Insufficient inventory")
		}

	}
	return nil, status.Errorf(codes.Unimplemented, "method Sell not implemented")
}

// func (UnimplementedInventoryServer) Reback(context.Context, *SellInfo) (*emptypb.Empty, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method Reback not implemented")
// }
