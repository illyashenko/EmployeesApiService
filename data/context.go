package data

import (
	"EmployeesApiService/configs"
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client  *redis.Client
	Context context.Context
}

func RedisContext(config configs.RedisConfig) RedisClient {
	return RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Address,
			Password: config.Password,
			DB:       config.DefaultDB,
		}),
		Context: context.Background(),
	}
}

func (receiver *RedisClient) Get(key string) *redis.StringCmd {
	return receiver.client.Get(receiver.Context, key)
}

func (receiver *RedisClient) Set(key string, value string, ttl time.Duration) *redis.StatusCmd {
	return receiver.client.Set(receiver.Context, key, value, ttl)
}

func (receiver *RedisClient) GetAllKeys() *redis.Cmd {
	return receiver.client.Do(receiver.Context, "KEYS", "*")
}

func (receiver *RedisClient) Subscribe(pattern string) *redis.PubSub {
	return receiver.client.PSubscribe(receiver.Context, pattern)
}
