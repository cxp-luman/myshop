package global

import (
	"myshop-api/goods_web/config"
	"myshop-api/goods_web/proto"

	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
)

var (
	NacosInfo *config.NacosInfo = &config.NacosInfo{}
	GoodsWebInfo *config.GoodsWebInfo = &config.GoodsWebInfo{}
	Trans ut.Translator
	Rdb *redis.Client
	Srvclient proto.GoodsClient
)