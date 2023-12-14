package eventService

import (
	"EmployeesApiService/appconfig"
	"EmployeesApiService/data/stores"
	"EmployeesApiService/mq"
)

type EventService struct {
	mqKafka mq.KafkaMq
	store   *stores.RedisStore
	config  appconfig.KafkaConfig
}

func NewEventService(store *stores.RedisStore, config appconfig.KafkaConfig) EventService {
	return EventService{
		store:   store,
		config:  config,
		mqKafka: mq.NewKafkaProducer(config),
	}
}

func (receiver EventService) Start() {
	receiver.store.SubscribeSetKeyEvents(receiver.mqKafka, receiver.config.Topics)
}
