package global

import (
	"go-boilerplate-api/configs"

	"github.com/Azure/azure-sdk-for-go/sdk/data/azcosmos"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Application struct {
	Config       configs.Configuration
	DB           *gorm.DB
	CosmosClient *azcosmos.Client
	Redis        *redis.Client
}

var App = new(Application)
