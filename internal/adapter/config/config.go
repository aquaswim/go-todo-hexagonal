package config

import "os"

type Config struct {
	DBConnection      string
	DBMigrationOnBoot bool
}

func New() *Config {
	return &Config{
		DBConnection:      os.Getenv("DB_CONNECTION"),
		DBMigrationOnBoot: os.Getenv("DB_MIGRATE_ON_BOOT") == "true",
	}
}
