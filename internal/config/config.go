package config

import (
	"os"
)

const configTag = "Config"

func LoadConfig() (*Config, error) {
	appConfig := APPConfig{
		Port: getEnv("APP_PORT", "5000"),
	}

	dbConfig := loadConfigDB()

	redisConfig := loadConfigRedis()

	config := &Config{
		DB:    dbConfig,
		REDIS: redisConfig,
		APP:   appConfig,
	}
	return config, nil
}

func loadConfigDB() DBConfig {
	dbConfig := DBConfig{
		DBName:   getEnv("DB_NAME", "kodein"),
		User:     getEnv("DB_USER", "kodein"),
		Password: getEnv("DB_PASSWORD", "developmentpass"),
		Host:     getEnv("DB_HOST", "127.0.0.1"),
		Port:     getEnv("DB_PORT", "5433"),
		SSLMode:  getEnv("DB_SSL_MODE", "disable"),
	}
	return dbConfig
}

func loadConfigRedis() RedisConfig {
	redisConfig := RedisConfig{
		Address:  getEnv("REDIS_ADDRESS", "localhost:6379"),
		Password: getEnv("REDIS_PASSWORD", ""),
		DB:       getEnv("REDIS_DB", "0"),
	}
	return redisConfig
}

func getEnv(key string, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
