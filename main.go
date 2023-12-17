package main

import (
	"EmployeesApiService/configs"
	"EmployeesApiService/data"
	"EmployeesApiService/services"
)

func main() {

	config := configs.NewConfig()
	context := data.RedisContext(config.RedisConfig)

	if config.UseRedisEvents {
		go services.NewEventService(&data.RedisStore{Redis: &context}, config.KafkaMq).Start()
	}

	services.NewServer(&data.RedisStore{Redis: &context}).Start()
}
