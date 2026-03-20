package pkg

import (
	"go-boilerplate-api/internal/pkg/cache"
	"go-boilerplate-api/internal/pkg/db"

	"github.com/google/wire"
)

var DataProviderSet = wire.NewSet(
	db.NewRDBMS,
	cache.NewRedis,
	cache.NewCacheClient,
	wire.Bind(new(cache.CacheClient), new(*cache.RedisCacheClient)),
)
