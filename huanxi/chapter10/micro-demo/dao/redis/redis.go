package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/rey/micro-demo/log"
)

var RedisClient *redis.Client

const (
	Addr = "localhost:6379"
)

func init() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: Addr,
	})

	if err := redis.Client.Ping(*RedisClient, context.Background()).Err(); err != nil {
		panic(err)
	}

	log.Logger.Info("redis init ok")

}
