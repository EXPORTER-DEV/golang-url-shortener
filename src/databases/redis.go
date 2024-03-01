package databases

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisDatabaseInterface interface {
	Get(key string) (value string, err error)
	Set(key string, value string) error
}

type RedisDatabase struct {
	client *redis.Client
}

func (r *RedisDatabase) Get(key string) (value string, err error) {
	return r.client.Get(key).Result()
}

func (r *RedisDatabase) Set(key string, value string) error {
	return r.client.Set(key, value, time.Duration(0)).Err()
}

func NewRedisDatabase(client *redis.Client) RedisDatabaseInterface {
	return &RedisDatabase{client}
}
