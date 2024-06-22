package config

import "os"

type DBConfig struct {
	Connection      string
	MigrationOnBoot bool
}

func DBConfigFromENV() *DBConfig {
	return &DBConfig{
		Connection:      os.Getenv("DB_CONNECTION"),
		MigrationOnBoot: os.Getenv("DB_MIGRATE_ON_BOOT") == "true",
	}
}
