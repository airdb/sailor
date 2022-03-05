package redisutil

import (
	"log"

	"github.com/go-redis/redis"
)

func NewClient(opt *redis.Options) *redis.Client {
	redisdb := redis.NewClient(opt)

	log.Println("redis status:", redisdb.Ping())

	return redisdb
}
