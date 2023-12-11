package dbcontext

import (
	"EmployeesApiService/appconfig"
	"EmployeesApiService/mq"
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisClient struct {
	client *redis.Client
	ctx    context.Context
}

func RedisContext(config appconfig.RedisConfig) RedisClient {
	return RedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     config.Address,
			Password: config.Password,
			DB:       config.DefaultDB,
		}),
		ctx: context.Background(),
	}
}

func (receiver *RedisClient) Get(key string) *redis.StringCmd {
	return receiver.client.Get(receiver.ctx, key)
}

func (receiver *RedisClient) Set(key string, value string, ttl time.Duration) *redis.StatusCmd {
	return receiver.client.Set(receiver.ctx, key, value, ttl)
}

func (receiver *RedisClient) GetAllKeys() *redis.Cmd {
	return receiver.client.Do(receiver.ctx, "KEYS", "*")
}

func (receiver *RedisClient) Subscribe(pattern string, kafkaMq mq.KafkaMq, topics []string) {
	sub := receiver.client.PSubscribe(receiver.ctx, pattern)
	for {
		messages, err := sub.Receive(receiver.ctx)
		if err != nil {
			break
		}
		if messages != nil {
			switch msg := messages.(type) {
			case *redis.Message:
				key := fmt.Sprintf("%v", msg.Payload)
				val, _ := receiver.Get(key).Result()

				for _, topic := range topics {
					message := kafka.Message{
						TopicPartition: kafka.TopicPartition{
							Topic:     &topic,
							Partition: kafka.PartitionAny,
						},
						Key:   []byte(key),
						Value: []byte(val),
					}
					kafkaMq.SendMassages(message)
				}

			}
		}
	}
}
