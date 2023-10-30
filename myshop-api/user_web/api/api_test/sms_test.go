package apitest_test

import (
	"fmt"
	"myshop-api/user_web/global"
	"testing"

	"github.com/go-redis/redis"
)

func TestSms(t *testing.T) {
	fmt.Println(global.UserWebInfo.RedisInfo.Host, global.UserWebInfo.RedisInfo.Port)
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", global.UserWebInfo.RedisInfo.Host, global.UserWebInfo.RedisInfo.Port),
		Password: global.UserWebInfo.RedisInfo.Password,
	})
	if rdb==nil{
		t.Error("failed")
	}
}