package rLock

import (
	"edit-your-project-name/data"
	"edit-your-project-name/slog"
	"edit-your-project-name/utils"
	"github.com/bwmarrin/snowflake"
	"strconv"
	"time"
)

// Lock 尝试获得并发锁
func Lock(uid uint64) (ok bool, lock snowflake.ID) {
	key := data.RDB.K("rLock", strconv.FormatUint(uid, 10))
	lock = utils.No.Generate()

	ok, err := data.RDB.SetNX(data.CTX, key, int64(lock), time.Minute).Result()
	if err != nil {
		slog.Err(slog.PS("rLock", uid), err)
		return false, lock
	}
	return ok, lock
}

const unlockLuaScript = `if redis.call("GET", KEYS[1]) == ARGV[1] then
        return redis.call("DEL", KEYS[1])
    else
        return 0
    end`

// Unlock 释放并发锁
func Unlock(uid uint64, lock snowflake.ID) (ok bool) {
	key := data.RDB.K("rLock", strconv.FormatUint(uid, 10))

	res, err := data.RDB.Eval(data.CTX, unlockLuaScript, []string{key}, int64(lock)).Result()
	if err != nil || res.(int64) == 0 {
		slog.Err(slog.PS("rLock", uid, lock), err)
		return false
	}
	return true
}
