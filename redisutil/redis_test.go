package redisutil_test

import (
	"testing"

	"github.com/airdb/sailor/redisutil"
	"github.com/go-redis/redis"
)

func TestNewRedisClient(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping testing in short mode")
	}

	opt := redis.Options{}
	opt.Addr = "127.0.0.1:6379"
	opt.DB = 2
	opt.Password = "airdb"

	redisdb := redisutil.NewClient(&opt)
	t.Log(redisdb.Ping())
}
