package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/aliirsyaadn/kodein/internal/config"
	"github.com/aliirsyaadn/kodein/internal/log"
	"github.com/go-redis/redis/v8"
)

const redisTag = "InternalRedisTag"

type RedisCache interface {
	Close()
	SetJSON(ctx context.Context, service string, id string, data interface{}, expired time.Duration) bool
	Get(ctx context.Context, service string, id string) string
	Del(ctx context.Context, service string, id string)
}

type redisCache struct {
	client *redis.Client
}

func NewClient(redisConfig config.RedisConfig) *redis.Client{
	db, err := strconv.Atoi(redisConfig.DB)
	if err != nil {
		log.ErrorDetail(redisTag, "error convert string to int: %v", err)
		return nil
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Address,
		Password: redisConfig.Password,
		DB:       db,
	})

	return rdb
}

func NewCacheWithClient(redisConfig config.RedisConfig) RedisCache {
	client := NewClient(redisConfig)
	return &redisCache{client}
}

func NewCache(client *redis.Client) RedisCache {
	return &redisCache{client}
}

func (r *redisCache) Close(){
	r.client.Close()
}

func (r *redisCache) SetJSON(ctx context.Context, service string, id string, data interface{}, expired time.Duration) bool {
	key := fmt.Sprintf("%s:%s", service, id)
	
	value, err := json.Marshal(data) 
	if err != nil {
		log.ErrorDetail(redisTag, "error parse data to json: %v", err)
		return false
	}

	r.client.Set(ctx, key, value, expired*time.Second)
	return true
}

func (r *redisCache) Get(ctx context.Context, service string, id string) string {
	key := fmt.Sprintf("%s:%s", service, id)

	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		log.ErrorDetail(redisTag, "error get data from redis: %v", err)
		return ""
	}

	return value
}

func (r *redisCache) Del(ctx context.Context, service string, id string) {
	key := fmt.Sprintf("%s:%s", service, id)
	r.client.Del(ctx, key)
}


