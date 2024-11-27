package db

import (
	"log"
	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

var RedisClient *redis.Client

// InitRedis initializes the Redis connection
func InitRedis() {
	// Redis running on localhost:6379
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379", // Use 'localhost' or the actual IP if running in Docker
		Password: "",
		DB:   0,                // Default DB
	})

	// Test the connection
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Unable to connect to Redis: %v", err)
	}

	log.Println("Successfully connected to Redis!")
}
