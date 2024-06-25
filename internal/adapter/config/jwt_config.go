package config

import (
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

type JwtConfig struct {
	Secret   string
	Duration time.Duration
}

func JwtConfigFromENV() *JwtConfig {
	log.Debug().Msg("load JWT config from ENV")

	duration, err := time.ParseDuration(os.Getenv("JWT_DURATION"))
	if err != nil {
		log.Warn().Msg("failed to parse JWT_DURATION, fallback to 1h")
		duration = time.Hour
	}

	return &JwtConfig{
		Secret:   os.Getenv("JWT_SECRET"),
		Duration: duration,
	}
}
