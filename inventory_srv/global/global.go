package global

import (
	"myshop/inventory_srv/config"

	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	InvSrvInfo config.InvSrvInfo
	NacosInfo   config.NacosInfo
)
