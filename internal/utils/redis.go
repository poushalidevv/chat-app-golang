package utils

import (
	"log"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

func NewRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v\n", err)
	}

	return rdb
}
