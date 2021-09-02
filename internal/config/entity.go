package config

type Config struct {
	DB    DBConfig
	REDIS RedisConfig
	APP   APPConfig
	PRODUCER NSQProducerConfig
	CONSUMER NSQConsumerConfig
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

type NSQProducerConfig struct {
	ServerConfig
}

type NSQConsumerConfig struct {
	ServerConfig
	MaxAttempts uint16
	MaxInFlight int
}

type ServerConfig struct {
	Host string
	Port string
}