package db

import (
	"context"
	"edit-your-project-name/conf"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func InitRDB() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Pwd,
		DB:       conf.Redis.DB,
	})

	if _, err := RDB.Ping(context.Background()).Result(); err != nil {
		conf.FatalExt("Redis ERROR", err)
	}
}

func K(keys ...string) string {
	key := conf.Redis.Prefix
	for _, k := range keys {
		key += conf.Redis.Sep + k
	}
	return key
}
