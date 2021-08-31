package config

type Config struct {
	DB    DBConfig
	REDIS RedisConfig
	APP   APPConfig
}

type APPConfig struct {
	Port string
}

type DBConfig struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}

type RedisConfig struct {
	Address  string
	Password string
	DB       string
}
