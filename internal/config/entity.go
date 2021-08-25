package config

type Config struct {
	DB  DBConfig
	APP APPConfig
}

type APPConfig struct {
	Port        string
}

type DBConfig struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
	SSLMode  string
}
