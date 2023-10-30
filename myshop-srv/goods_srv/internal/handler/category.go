package handler

import (
	"context"
	"encoding/json"
	"myshop/goods_srv/global"
	"myshop/goods_srv/internal/db"
	"myshop/goods_srv/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)


func (s *GoodsServer) GetAllCategorysList(ctx context.Context, req *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []db.Category
	global.DbClient.Where(&db.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}
func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}

	var category db.Category
	if result := global.DbClient.First(&category, req.Id); result.RowsAffected == 0{
		return nil, status.Errorf(codes.NotFound, "Product category does not exist")
	}

	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:       category.ID,
		Name: category.Name,
		Level: category.Level,
		IsTab: category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	var subCategorys []db.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	//preloads := "SubCategory"
	//if category.Level == 1 {
	//	preloads = "SubCategory.SubCategory"
	//}
	global.DbClient.Where(&db.Category{ParentCategoryID: req.Id}).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id: subCategory.ID,
			Name: subCategory.Name,
			Level: subCategory.Level,
			IsTab: subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}

	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}
func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := db.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	cMap["is_tab"] = req.IsTab
	if req.Level != 1 {
		//去查询父类目是否存在
		cMap["parent_category_id"] = req.ParentCategory
	}
	if tx := global.DbClient.Model(&db.Category{}).Create(cMap); tx.Error != nil {
		return nil, status.Errorf(codes.Internal, "mysql connect failed")
	}
	return &proto.CategoryInfoResponse{Id:int32(category.ID)}, nil
}
func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DbClient.Delete(&db.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &emptypb.Empty{}, nil
}
func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category db.Category
	if result := global.DbClient.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "Product category does not exist")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}
	global.DbClient.Save(&category)
	return &emptypb.Empty{}, nil
}