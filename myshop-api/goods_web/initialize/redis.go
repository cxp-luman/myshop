package initialize

import (
	"context"
	"fmt"
	"myshop-api/goods_web/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.GoodsWebInfo.RedisInfo.Host, global.GoodsWebInfo.RedisInfo.Port),
		Password: global.GoodsWebInfo.RedisInfo.Password,
	})
	_, err := rdb.Get(context.Background(), "test-key").Result()
	if err != nil {
		zap.S().Errorw("dial redis failed", "host", global.GoodsWebInfo.RedisInfo.Host, "port", global.GoodsWebInfo.RedisInfo.Password)
		panic("dial redis failed!")
	}
	global.Rdb = rdb
}
