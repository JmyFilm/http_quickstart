package data

import (
	"PROJECTNAME/conf"
	"PROJECTNAME/utils"
	"PROJECTNAME/xlog"
	"context"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
)

var RDB *rdbClient

func InitRDB() {
	RDB = initRDB(conf.Redis)

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

func initRDB(config conf.SRedis) *rdbClient {
	client := redis.NewClient(&redis.Options{
		Addr:        config.Addr,
		Password:    config.Pwd,
		DB:          config.DB,
		PoolSize:    config.PoolSize,
		DialTimeout: time.Duration(config.DialTimeout) * time.Second,
	})
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		xlog.Fatal(xlog.PS("Redis", config.Addr, "ERROR"), err)
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

	xlog.Info("Redis", config.Addr, "Time:", _rdbClient.Time(CTX).Val())
	return _rdbClient
}

func RLock(srcId string, ttl time.Duration) (ok bool, lock string) {
	key := RDB.K("lock", srcId)
	lock = utils.SnowFlakeId()

	ok, err := RDB.SetNX(CTX, key, lock, ttl).Result()
	if err != nil {
		xlog.Err(xlog.PS("RLock", srcId, "ERROR"), err)
		return false, lock
	}
	return ok, lock
}

const rUnlockLuaScript = `if redis.call("GET", KEYS[1]) == ARGV[1] then
	return redis.call("DEL", KEYS[1])
else
	return 0
end`

func RUnlock(srcId string, lock string) bool {
	key := RDB.K("lock", srcId)

	res, err := RDB.Eval(CTX, rUnlockLuaScript, []string{key}, lock).Result()
	affected, _ := res.(int64)
	if err != nil || affected == 0 {
		xlog.Err(xlog.PS("RUnlock", srcId, lock, "ERROR"), err)
		return false
	}
	return true
}
