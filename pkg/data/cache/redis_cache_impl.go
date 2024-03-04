package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"time"
)

type redisStoreImpl struct {
	ctx        *context.Context
	log        *log.Helper
	cacheStore *RedisStore
}

var IncrAndSetExpiryOnce = redis.NewScript(`
	local key = KEYS[1]
	local ttl = ARGV[1]
	
	local value = redis.call("INCR", key)
	if value == 1 then
		redis.call("EXPIRE", key, ttl)
	end
	return value
`)

func NewRedisRepositoryImpl(cacheStore *RedisStore, logger log.Logger) CacheInterface {
	return &redisStoreImpl{cacheStore: cacheStore, log: log.NewHelper(logger)}
}

func (rs *redisStoreImpl) Get(ctx context.Context, key string, value interface{}) error {
	bytes, err := rs.cacheStore.DbClient.Get(ctx, key).Bytes()
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, value)
	if err != nil {
		return err
	}
	return nil
}

func (rs *redisStoreImpl) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = rs.cacheStore.DbClient.Set(ctx, key, p, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rs *redisStoreImpl) SetMap(ctx context.Context, key string, values []interface{}, ttl int64) error {
	err := rs.cacheStore.DbClient.HMSet(ctx, key, values...).Err()
	if err != nil {
		fmt.Println("Error setting data in Redis:", err)
		return err
	}
	rs.ExpireKey(ctx, key, ttl)
	return nil
}

func (rs *redisStoreImpl) ExpireKey(ctx context.Context, key string, ttl int64) bool {
	res := rs.cacheStore.DbClient.Expire(ctx, key, time.Duration(ttl)*time.Second)
	return res.Val()
}

func (rs *redisStoreImpl) GetMap(ctx context.Context, key string) (map[string]string, error) {
	result, err := rs.cacheStore.DbClient.HGetAll(ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting data from Redis:", err)
		return nil, err
	}
	retrievedMap := make(map[string]string)
	for key, value := range result {
		retrievedMap[key] = value
	}
	return retrievedMap, nil
}

func (rs *redisStoreImpl) MultiSet(ctx context.Context, values []interface{}) error {
	res := rs.cacheStore.DbClient.MSet(ctx, values...)
	return res.Err()
}

func (rs *redisStoreImpl) MultiGet(ctx context.Context, keys []string) ([]interface{}, error) {
	res := rs.cacheStore.DbClient.MGet(ctx, keys...)
	return res.Val(), res.Err()
}

func (rs *redisStoreImpl) GetKeysByPrefix(ctx context.Context, prefix string) ([]string, error) {
	keys, err := rs.cacheStore.DbClient.Keys(ctx, prefix+"*").Result()
	return keys, err
}

func (rs *redisStoreImpl) IncrAndSetExpiryOnce(ctx context.Context, key string, ttl int64) (int64, error) {
	updatedValue, err := IncrAndSetExpiryOnce.Run(ctx, rs.cacheStore.DbClient, []string{key}, time.Duration(ttl)).Int64()
	return updatedValue, err
}

func (rs *redisStoreImpl) DecrementValue(ctx context.Context, key string) int64 {
	res := rs.cacheStore.DbClient.Decr(ctx, key)
	return res.Val()
}

func (rs *redisStoreImpl) IncrementBy(ctx context.Context, key string, incrValue int64) int64 {
	res := rs.cacheStore.DbClient.IncrBy(ctx, key, incrValue)
	return res.Val()
}

func (rs *redisStoreImpl) Delete(ctx context.Context, keys ...string) error {
	pipe := rs.cacheStore.DbClient.Pipeline()
	for _, key := range keys {
		pipe.Del(ctx, key)
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		rs.log.WithContext(ctx).Errorf("error while delete cache: %+v", err)
		return err
	}
	return nil
}
