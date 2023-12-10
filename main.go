package main

import (
	"EmployeesApiService/appconfig"
	"EmployeesApiService/data/dbcontext"
	"EmployeesApiService/data/stores"
	"EmployeesApiService/mq"
	"EmployeesApiService/services/httpServer"
)

func main() {

	config := appconfig.NewConfig()
	context := dbcontext.RedisContext(config.RedisConfig)

	go context.Subscribe("__key*__:set", mq.NewKafkaProducer(config.KafkaMq))

	server := httpServer.NewServer(&stores.RedisStore{Redis: &context})
	server.Start()
}
