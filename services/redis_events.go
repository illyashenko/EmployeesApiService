package services

import (
	"EmployeesApiService/configs"
	"EmployeesApiService/data"
	"EmployeesApiService/mq"
)

type EventService struct {
	mqKafka mq.KafkaMq
	store   *data.RedisStore
	config  configs.KafkaConfig
}

func NewEventService(store *data.RedisStore, config configs.KafkaConfig) EventService {
	return EventService{
		store:   store,
		config:  config,
		mqKafka: mq.NewKafkaProducer(config),
	}
}

func (receiver EventService) Start() {
	receiver.store.SubscribeSetKeyEvents(receiver.mqKafka, receiver.config.Topics)
}
