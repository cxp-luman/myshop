package handler

import (
	"context"
	"myshop/goods_srv/global"
	"myshop/goods_srv/internal/db"
	"myshop/goods_srv/internal/proto"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *GoodsServer) BannerList(ctx context.Context, req *emptypb.Empty) (*proto.BannerListResponse, error) {
	bannerListResponse := proto.BannerListResponse{}

	var banners []db.Banner
	result := global.DbClient.Find(&banners)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "storage connect error!")
	}
	bannerListResponse.Total = int32(result.RowsAffected)

	var bannerReponses []*proto.BannerResponse
	for _, banner := range banners {
		bannerReponses = append(bannerReponses, &proto.BannerResponse{
			Id:       banner.ID,
			Image:       banner.Image,
			Index: banner.Index,
			Url: banner.Url,
		})
	}
	zap.S().Info("get banner select!")
	bannerListResponse.Data = bannerReponses

	return &bannerListResponse, nil
}
func (s *GoodsServer) CreateBanner(ctx context.Context, req *proto.BannerRequest) (*proto.BannerResponse, error) {
	banner := db.Banner{}

	banner.Image = req.Image
	banner.Index = req.Index
	banner.Url = req.Url

	global.DbClient.Save(&banner)

	return &proto.BannerResponse{Id:banner.ID}, nil
}
func (s *GoodsServer) DeleteBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	if result := global.DbClient.Delete(&db.Banner{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) UpdateBanner(ctx context.Context, req *proto.BannerRequest) (*emptypb.Empty, error) {
	var banner db.Banner

	if result := global.DbClient.First(&banner, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "轮播图不存在")
	}

	if req.Url != "" {
		banner.Url = req.Url
	}
	if req.Image != "" {
		banner.Image = req.Image
	}
	if req.Index != 0 {
		banner.Index = req.Index
	}

	global.DbClient.Save(&banner)

	return &emptypb.Empty{}, nil
}