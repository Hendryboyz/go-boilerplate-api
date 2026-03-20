package configs

type Configuration struct {
	Server  ServerConfig `mapstructure:"server" yaml:"server"`
	Session Session      `mapstructure:"session" yaml:"session"`
	DB      DBConfig     `mapstructure:"db" yaml:"db"`
	Cosmos  CosmosConfig `mapstructure:"cosmos" yaml:"cosmos"`
	Redis   RedisConfig  `mapstructure:"redis" yaml:"redis"`
	Domain  Domain       `mapstructure:"domain" yaml:"domain"`
	Queue   QueueConfig  `mapstructure:"queue" yaml:"queue"`
}

type ServerConfig struct {
	ModuleCode      string `mapstructure:"moduleCode" yaml:"moduleCode"`
	ModuleShortName string `mapstructure:"moduleShortName" yaml:"moduleShortName"`
	HttpPort        int    `mapstructure:"httpPort" yaml:"httpPort"`
	GrpcPort        int    `mapstructure:"grpcPort" yaml:"grpcPort"`
	Profiling       bool   `mapstructure:"profiling" yaml:"profiling"`
	Environment     string `mapstructure:"environment" yaml:"environment"`
	AllowOrigins    string `mapstructure:"allowOrigins" yaml:"allowOrigins"`
}

type Session struct {
	TokenName string `mapstructure:"tokenName" yaml:"tokenName"`
}

type DBConfig struct {
	DSN             string `mapstructure:"dsn" yaml:"dsn"`
	MaxConnection   int    `mapstructure:"maxConnection" yaml:"maxConnection"`
	CreateBatchSize int    `mapstructure:"createBatchSize" yaml:"createBatchSize"`
	AutoMigration   bool   `mapstructure:"autoMigration" yaml:"autoMigration"`
}

type CosmosConfig struct {
	CosmosDatabaseURL string `mapstructure:"cosmosDatabaseURL" yaml:"cosmosDatabaseURL"`
	CosmosDatabaseKey string `mapstructure:"cosmosDatabaseKey" yaml:"cosmosDatabaseKey"`
}

type RedisConfig struct {
	URL            string `mapstructure:"url" yaml:"url"`
	DB             int    `mapstructure:"db" yaml:"db"`
	PoolSize       int    `mapstructure:"poolSize" yaml:"poolSize"`
	Password       string `mapstructure:"password" yaml:"password"`
	CachedPrefix   string `mapstructure:"cachedPrefix" yaml:"cachedPrefix"`
	DefaultTTLSecs int    `mapstructure:"defaultTTLSecs" yaml:"defaultTTLSecs"`
}

type Domain struct {
	Bill BillConfig `mapstructure:"bill" yaml:"bill"`
}

type BillConfig struct {
	ExpirationYear  int    `mapstructure:"expirationYear" yaml:"expirationYear"`
	UploadDirectory string `mapstructure:"uploadDirectory" yaml:"uploadDirectory"`
}

type QueueConfig struct {
	ServiceBus    ServiceBusConfig `mapstructure:"serviceBus" yaml:"serviceBus"`
	BillQueueName string           `mapstructure:"billQueueName" yaml:"billQueueName"`
}

type ServiceBusConfig struct {
	HostName      string `mapstructure:"hostName" yaml:"hostName"`
	AccessKeyName string `mapstructure:"accessKeyName" yaml:"accessKeyName"`
	AccessKey     string `mapstructure:"accessKey" yaml:"accessKey"`
}
