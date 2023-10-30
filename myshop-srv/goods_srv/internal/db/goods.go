package db

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(100);not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"default:false;not null"`
}

type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(100);not null"`
	Logo string `gorm:"type:varchar(200);not null;default:''"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category
	BrandsId int32 `gorm:"type:int;index:idx_category_brand,unique"`
	*Brands
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel

	CategoryID int32 `gorm:"type:int;not null"`
	Category Category
	BrandsID int32 `gorm:"type:int;not null"`
	Brands Brands

	OnSale bool `gorm:"default:false;not null"`
	ShipFree bool `gorm:"default:false;not null"`
	IsNew bool `gorm:"default:false;not null"`
	IsHot bool `gorm:"default:false;not null"`

	Name  string `gorm:"type:varchar(50);not null"`
	GoodsSn string `gorm:"type:varchar(50);not null"`
	ClickNum int32 `gorm:"type:int;default:0;not null"`
	SoldNum int32 `gorm:"type:int;default:0;not null"`
	FavNum int32 `gorm:"type:int;default:0;not null"`
	MarketPrice float32 `gorm:"not null"`
	ShopPrice float32 `gorm:"not null"`
	GoodsBrief string `gorm:"type:varchar(100);not null"`
	Images GormList `gorm:"type:varchar(1000);not null"` // 解决存储多张图片的问题：1、另建一张表（关联查询影响性能）；2、自定义类型存储在当前表的一个字段中
	DescImages GormList `gorm:"type:varchar(1000);not null"`
	GoodsFrontImage string `gorm:"type:varchar(200);not null"`
}