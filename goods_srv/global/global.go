package global

import (
	"myshop/goods_srv/config"

	"gorm.io/gorm"
)

var (
	DbClient          *gorm.DB
	GoodsSrvInfo config.GoodSrvInfo
	NacosInfo   config.NacosInfo
)
