package handler

import (
	"context"
	"myshop/goods_srv/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

// func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
// 	brandListClient := global.DbClient
// 	if req.PriceMin > 0 {
// 		brandListClient = brandListClient.Where("shop_price > ?", req.PriceMin)
// 	}
// 	if req.PriceMax > 0 {
// 		brandListClient = brandListClient.Where("shop_price < ?", req.PriceMax) 
// 	}
// 	if req.IsHot {
// 		brandListClient = brandListClient.Where(&db.Goods{IsHot: true})
// 	}
// 	if req.IsNew {
// 		brandListClient = brandListClient.Where(&db.Goods{IsNew: true})
// 	}
// 	if req.KeyWords != "" {
// 		brandListClient = brandListClient.Where("name like ? or goods_brief like ?", "%"+req.KeyWords+"%", "%"+req.KeyWords+"%")
// 	}
// 	if req.Brand > 0 {
// 		brandListClient = brandListClient.Where("brands_id", req.Brand)
// 	}
// 	goods := []db.Goods{}
// 	result := brandListClient.Scopes(utils.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&goods)
// 	if result.Error != nil {
// 		return nil, status.Errorf(codes.Internal, "db connect failed")
// 	}
// 	if req.TopCategory == 0 {
// 		global.DbClient.Where("select * from ")
// 	}
// }
func (s *GoodsServer) BatchGetGoods(ctx context.Context, req *proto.BatchGoodsIdInfo) (*proto.GoodsListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetGoods not implemented")
}
// func (s *GoodsServer) CreateGoods(ctx context.Context, *CreateGoodsInfo) (*GoodsInfoResponse, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method CreateGoods not implemented")
// }
// func (s *GoodsServer) DeleteGoods(ctx context.Context, *DeleteGoodsInfo) (*emptypb.Empty, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method DeleteGoods not implemented")
// }
// func (s *GoodsServer) UpdateGoods(ctx context.Context, *CreateGoodsInfo) (*emptypb.Empty, error) {
// 	return nil, status.Errorf(codes.Unimplemented, "method UpdateGoods not implemented")
// }