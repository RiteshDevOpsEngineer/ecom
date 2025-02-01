package ports

import (
	"context"
	"time"
)

type RedisClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	// Del(ctx context.Context, keys ...string) error
	// HGet(ctx context.Context, key, field string) (string, error)
	// HSet(ctx context.Context, key, field string, value interface{}) error
	// Exists(ctx context.Context, key string) (bool, error)
	// Incr(ctx context.Context, key string) (int64, error)
	// Decr(ctx context.Context, key string) (int64, error)
	// Expire(ctx context.Context, key string, expiration time.Duration) (bool, error)
}
