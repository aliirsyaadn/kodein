package config

import (
	"os"
)

const configTag = "Config"

func LoadConfig() (*Config, error) {
	appConfig := &APPConfig{
		Port:        getEnv("APP_PORT", "5000"),
	}

	dbConfig := loadConfigDB()

	config := &Config{
		DB:  *dbConfig,
		APP: *appConfig,
	}
	return config, nil
}

func loadConfigDB() *DBConfig {
	dbConfig := &DBConfig{
		DBName:   getEnv("DB_NAME", "kodein"),
		User:     getEnv("DB_USER", "kodein"),
		Password: getEnv("DB_PASSWORD", "developmentpass"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "5432"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}
	return dbConfig
}

func getEnv(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
