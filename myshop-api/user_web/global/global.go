package global

import (
	"myshop-api/user_web/config"
	"myshop-api/user_web/proto"

	ut "github.com/go-playground/universal-translator"
	"github.com/redis/go-redis/v9"
)

var (
	UserWebInfo *config.UserWebInfo = &config.UserWebInfo{}
	Trans ut.Translator
	Rdb *redis.Client
	Srvclient proto.UserClient
)