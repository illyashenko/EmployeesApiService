package stores

import (
	"EmployeesApiService/data/dbcontext"
	"EmployeesApiService/models"
	"EmployeesApiService/mq"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/redis/go-redis/v9"
	"reflect"
)

type RedisStore struct {
	Redis *dbcontext.RedisClient
}

func (receiver *RedisStore) Employee(key string) (models.Employee, error) {
	value, err := receiver.Redis.Get(key).Result()
	if err != nil {
		return models.Employee{}, err
	}
	var employee models.Employee
	errJson := json.Unmarshal([]byte(value), &employee)
	return employee, errJson
}

func (receiver *RedisStore) Employees() ([]models.Employee, error) {
	value, err := receiver.Redis.GetAllKeys().Slice()
	if err != nil {
		return nil, err
	}
	var employees []models.Employee
	if reflect.TypeOf(value).Kind() == reflect.Slice {
		for key := range value {
			employee, errJson := receiver.Employee(fmt.Sprintf("%v", value[key]))

			if errJson != nil {
				continue
			}
			employees = append(employees, employee)
		}
	}
	return employees, err
}

func (receiver *RedisStore) SubscribeSetKeyEvents(mqKafka mq.KafkaMq, topics []string) {

	subscriber := receiver.Redis.Subscribe("__key*__:set")

	for {
		messages, err := subscriber.Receive(receiver.Redis.Context)
		if err != nil {
			break
		}
		if messages != nil {
			switch msg := messages.(type) {
			case *redis.Message:
				key := fmt.Sprintf("%v", msg.Payload)
				val, _ := receiver.Redis.Get(key).Result()

				for _, topic := range topics {
					message := kafka.Message{
						TopicPartition: kafka.TopicPartition{
							Topic:     &topic,
							Partition: kafka.PartitionAny,
						},
						Key:   []byte(key),
						Value: []byte(val),
					}
					mqKafka.SendMassages(message)
				}

			}
		}
	}
}
