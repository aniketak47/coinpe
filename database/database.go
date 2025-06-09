package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func New(cfg DBConfiguration, options ...func(*DBConnOptions)) (*gorm.DB, error) {
	conn := &DBConnOptions{
		migrations: []interface{}{},
		logConfig: DBLogConfig{
			MigrationsLogLevel: logger.Error,
			DefaultLogLevel:    logger.Error,
		},
	}
	for _, o := range options {
		o(conn)
	}

	gormDB, err := gorm.Open(postgres.Open(getDSN(cfg)), &gorm.Config{
		Logger: logger.Default.LogMode(conn.logConfig.DefaultLogLevel),
	})
	if err != nil {
		log.Fatal("database connection error", err.Error())
		return nil, err
	}

	if conn.replicaConfig != nil {
		gormDB.Use(dbresolver.Register(dbresolver.Config{
			Replicas: []gorm.Dialector{
				postgres.Open(getDSN(*conn.replicaConfig)),
			},
			Policy: dbresolver.RandomPolicy{},
		}))
	}

	if len(conn.migrations) > 0 {
		err = GetSilentSession(gormDB).AutoMigrate(conn.migrations...)
		if err != nil {
			log.Fatalf("database migration error: %s", err)
		}
	}

	return gormDB, nil
}

func WithLogConfig(cfg DBLogConfig) func(*DBConnOptions) {
	return func(db *DBConnOptions) {
		db.logConfig = cfg
	}
}

func WithMigrations(migrations []interface{}) func(*DBConnOptions) {
	return func(db *DBConnOptions) {
		db.migrations = migrations
	}
}

func getDSN(cfg DBConfiguration) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)
}

func GetSilentSession(db *gorm.DB) *gorm.DB {
	return db.Session(&gorm.Session{Logger: logger.Default.LogMode(logger.Silent)})
}
