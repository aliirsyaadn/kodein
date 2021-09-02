package nsq

import (
	"fmt"

	"github.com/aliirsyaadn/kodein/internal/config"
)

//DSN FORMAT : 127.0.0.1:4150
func ParseDSN(cfg config.ServerConfig) string {
	return fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
}