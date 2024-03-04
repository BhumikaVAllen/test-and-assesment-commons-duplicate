package cache

import (
	"context"
	"time"
)

// CacheInterface Cache repository
type CacheInterface interface {
	Set(ctx context.Context, key string, src interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, dest interface{}) error
	ExpireKey(ctx context.Context, key string, ttl int64) bool
	SetMap(ctx context.Context, key string, values []interface{}, ttl int64) error
	GetMap(ctx context.Context, key string) (map[string]string, error)
	IncrAndSetExpiryOnce(ctx context.Context, key string, ttl int64) (int64, error)
	DecrementValue(ctx context.Context, key string) int64
	IncrementBy(ctx context.Context, key string, incrValue int64) int64
	MultiSet(ctx context.Context, values []interface{}) error
	MultiGet(ctx context.Context, keys []string) ([]interface{}, error)
	GetKeysByPrefix(ctx context.Context, prefix string) ([]string, error)
	Delete(ctx context.Context, keys ...string) error
}
