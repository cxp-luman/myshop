package initialize

import (
	"context"
	"fmt"
	"myshop-api/user_web/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedisClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", global.UserWebInfo.RedisInfo.Host, global.UserWebInfo.RedisInfo.Port),
		Password: global.UserWebInfo.RedisInfo.Password,
	})
	_, err := rdb.Get(context.Background(), "test-key").Result()
	if err != nil {
		zap.S().Errorw("dial redis failed", "host", global.UserWebInfo.RedisInfo.Host, "port", global.UserWebInfo.RedisInfo.Password)
		panic("dial redis failed!")
	}
	global.Rdb = rdb
}
