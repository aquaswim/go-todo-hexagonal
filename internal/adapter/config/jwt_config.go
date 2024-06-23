package config

import (
	"os"
	"time"
)

type JwtConfig struct {
	Secret   string
	Duration time.Duration
}

func JwtConfigFromENV() *JwtConfig {
	duration, err := time.ParseDuration(os.Getenv("JWT_DURATION"))
	if err != nil {
		//todo: log something to console
		duration = time.Hour
	}
	return &JwtConfig{
		Secret:   os.Getenv("JWT_SECRET"),
		Duration: duration,
	}
}
