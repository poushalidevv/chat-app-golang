package db

import (
    "context"
    "fmt"
    "log"
    "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr: "redis:6379", // Redis server address
    })

    _, err := RedisClient.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Unable to connect to Redis: %v", err)
    }

    fmt.Println("Successfully connected to Redis!")
}
