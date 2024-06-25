package config

import (
	"github.com/rs/zerolog/log"
	"os"
)

type DBConfig struct {
	Connection      string
	MigrationOnBoot bool
}

func DBConfigFromENV() *DBConfig {
	log.Debug().Msg("load dbconfig from environment")

	return &DBConfig{
		Connection:      os.Getenv("DB_CONNECTION"),
		MigrationOnBoot: os.Getenv("DB_MIGRATE_ON_BOOT") == "true",
	}
}
