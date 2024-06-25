package config

import (
	"github.com/rs/zerolog/log"
	"os"
)

type RestConfig struct {
	Port string
}

func RestConfigFromENV() *RestConfig {
	log.Debug().Msg("load rest config from environment")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &RestConfig{
		Port: port,
	}
}
