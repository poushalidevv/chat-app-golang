package config

import (
	"log"

	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ServerAddress string
	PostgresDSN   string
	RedisAddress  string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		PostgresDSN:   getEnv("POSTGRES_DSN", "user=postgres password=pass dbname=chat_app sslmode=disable"),
		RedisAddress:  getEnv("REDIS_ADDRESS", "localhost:6379"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
