package database

import "gorm.io/gorm/logger"

type DBConfiguration struct {
	Name     string `env:"NAME"`
	Username string `env:"USER"`
	Password string `env:"PASSWORD"`
	Host     string `env:"HOST"`
	Port     string `env:"PORT"`
	SSLMode  string `env:"SSL_MODE"`
	LogMode  bool   `env:"LOG_MODE"`
}

type DBLogConfig struct {
	DefaultLogLevel    logger.LogLevel
	MigrationsLogLevel logger.LogLevel
}

type DBConnOptions struct {
	migrations    []interface{}
	logConfig     DBLogConfig
	replicaConfig *DBConfiguration
}
