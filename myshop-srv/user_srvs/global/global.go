package global

import (
	"myshop/user_srvs/config"

	"gorm.io/gorm"
)

var (
	DB          *gorm.DB
	UserSrvInfo config.UserSrvInfo
	NacosInfo   config.NacosInfo
)
