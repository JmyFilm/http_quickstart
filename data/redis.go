package data

import (
	"context"
	"edit-your-project-name/config"
	"edit-your-project-name/slog"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

var RDB *rdbClient

func InitRDB() {
	RDB = initRDB(config.Redis)

	for _, fn := range waitRDBFn {
		fn()
	}
}

// ====^

var waitRDBFn []func()
var CTX = context.Background()

type rdbClient struct {
	*redis.Client
	K func(keys ...string) string
}

func WaitRDBExec(fn func()) {
	waitRDBFn = append(waitRDBFn, fn)
}

func D(d time.Duration) time.Duration {
	return d * 24 * time.Hour
}

func H(h time.Duration) time.Duration {
	return h * time.Hour
}

func M(m time.Duration) time.Duration {
	return m * time.Minute
}

func S(s time.Duration) time.Duration {
	return s * time.Second
}

func initRDB(config config.RedisConfig) *rdbClient {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Pwd,
		DB:       config.DB,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		slog.Fatal(slog.PS("Redis", config.Addr, "ERROR"), err)
	}

	var _rdbClient = &rdbClient{
		Client: client,
	}
	_rdbClient.K = func(keys ...string) string {
		if config.Prefix != "" {
			keys = append([]string{config.Prefix}, keys...)
		}
		return strings.Join(keys, config.Sep)
	}

	return _rdbClient
}
