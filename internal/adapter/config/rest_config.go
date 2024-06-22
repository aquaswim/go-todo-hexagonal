package config

import "os"

type RestConfig struct {
	Port string
	DB   *DBConfig
}

func RestConfigFromENV() *RestConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &RestConfig{
		Port: port,
		DB:   DBConfigFromENV(),
	}
}