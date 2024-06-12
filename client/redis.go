package client

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Redis(redisURL string) *redis.Client {
	opts, err := redis.ParseURL(redisURL)
	if err != nil {
		panic(err)
	}

	return redis.NewClient(opts)
}

func RedisConnection(rdb *redis.Client) bool {
	_, err := rdb.Ping(context.Background()).Result()
	return err == nil
}
