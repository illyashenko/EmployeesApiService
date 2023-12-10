package mq

import (
	"EmployeesApiService/appconfig"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

type KafkaMq struct {
	Producer *kafka.Producer
}

func NewKafkaProducer(config appconfig.KafkaConfig) KafkaMq {

	producer, err := kafka.NewProducer(configure(config))

	if err != nil {
		log.Fatal("Fail create KafkaConfig producer ")
	}

	return KafkaMq{Producer: producer}
}

func configure(config appconfig.KafkaConfig) *kafka.ConfigMap {

	kafkaConfig := kafka.ConfigMap{}

	for _, element := range config.Config {
		for key, value := range element {
			_ = kafkaConfig.SetKey(key, value)
		}
	}
	return &kafkaConfig
}

func (kafkaMq KafkaMq) SendMassages(message kafka.Message) {

	delChan := make(chan kafka.Event)

	err := kafkaMq.Producer.Produce(&message, delChan)

	if err != nil {
		log.Println(err.Error())
	}
	answer := <-delChan
	msg := answer.(*kafka.Message)

	if msg.TopicPartition.Error != nil {
		log.Println(msg.TopicPartition.Error.Error())
	}

	close(delChan)
}
