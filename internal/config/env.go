package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	DevDatabaseName  string
	ProdDatabaseName string
	TestDatabaseName string
	ApiHost          string
	ApiPort          string
	TrustedProxies   []string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	cfg := &Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		DevDatabaseName:  os.Getenv("DEV_DATABASE_NAME"),
		ProdDatabaseName: os.Getenv("PROD_DATABASE_NAME"),
		TestDatabaseName: os.Getenv("TEST_DATABASE_NAME"),
		ApiHost:          os.Getenv("API_HOST"),
		ApiPort:          os.Getenv("API_PORT"),
	}

	proxies := os.Getenv("TRUSTED_PROXIES")
	if proxies != "" {
		cfg.TrustedProxies = strings.Split(proxies, ",")
	} else {
		cfg.TrustedProxies = []string{}
	}

	return cfg, nil
}
