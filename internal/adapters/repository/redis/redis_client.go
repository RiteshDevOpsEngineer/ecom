package redis

import (
	"context"
	"time"

	"github.com/RiteshDevOpsEngineer/ecom/internal/adapters/database"
	"github.com/RiteshDevOpsEngineer/ecom/internal/core/ports"
)

type GoRedisClient struct{}

func NewGoRedisClient() ports.RedisClient {
	return &GoRedisClient{}
}

func (r *GoRedisClient) Get(ctx context.Context, key string) (string, error) {
	client := database.GetRedisClient()
	val, err := client.Get(ctx, key).Result()
	return val, err
}

func (r *GoRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	client := database.GetRedisClient()
	return client.Set(ctx, key, value, expiration).Err()
}

// func (r *GoRedisClient) Del(ctx context.Context, keys ...string) error {
// 	cmd := r.Client.Del(ctx, keys...)
// 	_, err := cmd.Result()
// 	return err
// }

// func (r *GoRedisClient) HGet(ctx context.Context, key, field string) (string, error) {
// 	return r.Client.HGet(ctx, key, field).Result()
// }

// func (r *GoRedisClient) HSet(ctx context.Context, key, field string, value interface{}) error {
// 	return r.Client.HSet(ctx, key, field, value).Err()
// }

// func (r *GoRedisClient) Incr(ctx context.Context, key string) (int64, error) {
// 	return r.Client.Incr(ctx, key).Result()
// }

// func (r *GoRedisClient) Exists(ctx context.Context, key string) (bool, error) {
// 	exists, err := r.Client.Exists(ctx, key).Result()
// 	return exists == 1, err
// }

// func (r *GoRedisClient) Expire(ctx context.Context, key string, expiration time.Duration) (bool, error) {
// 	return r.Client.Expire(ctx, key, expiration).Result()
// }
