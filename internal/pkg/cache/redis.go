package cache

import (
	"context"
	"crypto/tls"
	"go-boilerplate-api/global"
	"go-boilerplate-api/internal/pkg/log"

	"github.com/redis/go-redis/v9"
)

func getRedisOptions() *redis.Options {
	redisOptions := &redis.Options{
		Addr:     global.App.Config.Redis.URL,
		Password: global.App.Config.Redis.Password,
		DB:       global.App.Config.Redis.DB,
		PoolSize: global.App.Config.Redis.PoolSize,
	}
	if global.App.Config.Server.Environment == "production" {
		redisOptions.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
	}
	return redisOptions
}

func NewRedis() *redis.Client {
	redisOptions := getRedisOptions()
	rdb := redis.NewClient(redisOptions)

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Error("failed to connect to redis", log.String("reason", err.Error()))
	}
	return rdb
}
