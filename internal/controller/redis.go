package controller

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func RedisGet(rdb *redis.Client, key string) string {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}
	return val
}
