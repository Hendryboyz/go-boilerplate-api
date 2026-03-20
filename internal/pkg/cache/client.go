package cache

import (
	"context"
	"encoding/json"
	"go-boilerplate-api/global"
	"go-boilerplate-api/internal/pkg/log"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheClient interface {
	GetFromCache(key string) ([]byte, error)
	SetToCacheDefaultTTL(key string, value any) error
	SetToCache(key string, value any, ttl time.Duration) error
}

type RedisCacheClient struct {
	instance *redis.Client
}

func NewCacheClient(instance *redis.Client) *RedisCacheClient {
	return &RedisCacheClient{
		instance: instance,
	}
}

var ctx = context.Background()

func (r *RedisCacheClient) GetFromCache(key string) ([]byte, error) {
	keyCount, _ := r.instance.Exists(ctx, key).Result()
	if keyCount == 0 {
		return nil, nil
	}
	rawCachedValue, err := r.instance.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(rawCachedValue), nil
}

func (r *RedisCacheClient) SetToCacheDefaultTTL(key string, value any) error {
	log.Debug("set to cache", log.String("key", key))
	defaultTTLSecs := global.App.Config.Redis.DefaultTTLSecs * int(time.Second)
	return r.SetToCache(key, value, time.Duration(defaultTTLSecs))
}

func (r *RedisCacheClient) SetToCache(key string, value any, ttl time.Duration) error {
	bytesValue, _ := json.Marshal(value)
	jsonString := string(bytesValue)
	return r.instance.Set(ctx, key, jsonString, ttl).Err()
}
