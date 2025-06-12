package config

import (
	"coinpe/database"
	"coinpe/pkg/constants"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config Config
}

type Config struct {
	Environment        constants.AppEnv         `env:"ENVIRONMENT"`
	Debug              bool                     `env:"DEBUG"`
	Server             ServerConfiguration      `env:",prefix=SERVER_"`
	MainDatabase       database.DBConfiguration `env:",prefix=MAIN_DB_"`
	ReaderDB           database.DBConfiguration `env:",prefix=READER_DB_"`
	RedisConfiguration RedisConfiguration       `env:",prefix=REDIS_"`
	FeatureFlags       string                   `env:"FEATURE_FLAGS"`
	VPCProxyCIDR       string                   `env:"VPC_PROXY_CIDR"`
	JWTConfiguration   JWTConfiguration
}

type ServerConfiguration struct {
	Port      string `env:"PORT"`
	GRPCPort  string `env:"GRPC_PORT"`
	PublicURL string `env:"PUBLIC_URL"`
}

type RedisConfiguration struct {
	RedisConnectionAddress string `env:"CONNECTION_ADDRESS"`
	RedisPassword          string `env:"PASSWORD"`
}

type JWTConfiguration struct {
	SecretKey                             string `env:"JWT_SECRET_KEY"`
	PartialAuthAccessTokenExpiryInSeconds int    `env:"PARTIAL_AUTH_ACCESS_TOKEN_EXPIRY_IN_SECONDS"`
	FullAuthAccessTokenExpiryInSeconds    int    `env:"FULL_AUTH_ACCESS_TOKEN_EXPIRY_IN_SECONDS"`
	FullAuthRefreshTokenExpiryInSeconds   int    `env:"FULL_AUTH_REFRESH_TOKEN_EXPIRY_IN_SECONDS"`
}
