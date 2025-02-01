package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
	onceRedis     sync.Once
	redisClient   *redis.Client
)

func InitializeRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}

}

func GetRedisClient() *redis.Client {
	onceRedis.Do(func() {
		InitializeRedis()
	})
	return redisClient
}
