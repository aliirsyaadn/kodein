package redis

import (
	"strconv"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/go-redis/redis/v8"
)

const intRedisTag = "InternalRedisTag"

func ConnectRedis(redisConfig config.RedisConfig) *redis.Client {
	db, err := strconv.Atoi(redisConfig.DB)
	if err != nil {
		log.ErrorDetail(intRedisTag, "error convert string to int: %v", err)
		return nil
	}
	
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       db,
	})

	return rdb
}