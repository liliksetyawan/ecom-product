package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	_ "log"
)

var AppConfig *Config

type Config struct {
	DBHost     string `envconfig:"DB_HOST" required:"true"`
	DBPort     int    `envconfig:"DB_PORT" required:"true"`
	DBUser     string `envconfig:"DB_USER" required:"true"`
	DBPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DBName     string `envconfig:"DB_NAME" required:"true"`

	JWTSecret string `envconfig:"JWT_SECRET" required:"true"`
	JWTExpiry int    `envconfig:"JWT_EXPIRY" default:"72"`
}

func LoadConfig() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	AppConfig = &cfg
}
