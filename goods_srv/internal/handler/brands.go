package handler

import (
	"context"
	"myshop/goods_srv/global"
	"myshop/goods_srv/internal/db"
	"myshop/goods_srv/internal/proto"
	"myshop/goods_srv/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	dbClient := global.DbClient
	brands := []db.Brands{}
	var total int64
	result := dbClient.Scopes(utils.Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)
	if result.Error != nil {
		zap.S().Fatal("filter find brand list failed")
		return &proto.BrandListResponse{}, result.Error
	}
	global.DbClient.Model(&brands).Count(&total)
	brandsInfo := []*proto.BrandInfoResponse{}
	for _, brand := range brands {
		brandsInfo = append(brandsInfo, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	zap.S().Info("call brand list")
	return &proto.BrandListResponse{
		Total: int32(total),
		Data:  brandsInfo,
	}, nil
}

func (s *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	brands := db.Brands{}
	if result := global.DbClient.Select("id").Where("name = ?", req.Name).Find(&brands); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "brand is existed")
	}
	brands.ID = req.Id
	brands.Logo = req.Logo
	brands.Name = req.Name
	result := global.DbClient.Create(&brands)
	if result.Error != nil {
		return nil, result.Error
	}
	return &proto.BrandInfoResponse{
		Id:   brands.ID,
		Name: brands.Name,
		Logo: brands.Logo}, nil
}

func (s *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DbClient.Delete(&db.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "brands is not existed")
	}
	return &emptypb.Empty{}, nil
}
// todo: modify
func (s *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if req.Name == "" && req.Logo == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name and logo fields cannot all be null values")
	}
	brands := db.Brands{}
	if result := global.DbClient.Where("id = ?", req.Id).First(&brands); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "brand is not existed")
	}
	if req.Name != "" {
		brands.Name = req.Name
	}
	if req.Logo != "" {
		brands.Logo = req.Logo
	}
	if result := global.DbClient.Model(&db.Brands{}).Updates(map[string]interface{}{"name":brands.Name, "logo":brands.Logo}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "brand is not existed")
	}
	return &emptypb.Empty{}, nil
}
