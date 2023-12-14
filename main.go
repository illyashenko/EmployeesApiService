package main

import (
	"EmployeesApiService/appconfig"
	"EmployeesApiService/data/dbcontext"
	"EmployeesApiService/data/stores"
	"EmployeesApiService/services/eventService"
	"EmployeesApiService/services/httpServer"
)

func main() {

	config := appconfig.NewConfig()
	context := dbcontext.RedisContext(config.RedisConfig)

	if config.UseRedisEvents {
		go eventService.NewEventService(&stores.RedisStore{Redis: &context}, config.KafkaMq).Start()
	}

	httpServer.NewServer(&stores.RedisStore{Redis: &context}).Start()
}
