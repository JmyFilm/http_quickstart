package db

import (
	"context"
	"edit-your-project-name/conf"
	"github.com/redis/go-redis/v9"
)

var RDB RDbClient

func InitRDB() {
	RDB = initRDB(conf.Redis)
}

type RDbClient struct {
	*redis.Client
	K func(keys ...string) string
}

func initRDB(config conf.RedisConfig) RDbClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pwd,
		DB:       config.DB,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		conf.FatalExt("Redis ERROR", err)
	}
	return RDbClient{
		Client: client,
		K: func(keys ...string) string {
			key := config.Prefix
			for _, k := range keys {
				key += config.Sep + k
			}
			return key
		},
	}
}
