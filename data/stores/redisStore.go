package stores

import (
	"EmployeesApiService/data/dbcontext"
	"EmployeesApiService/models"
	"encoding/json"
	"fmt"
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
		for key, _ := range value {
			employee, errJson := receiver.Employee(fmt.Sprintf("%v", value[key]))

			if errJson != nil {
				continue
			}
			employees = append(employees, employee)
		}
	}
	return employees, err
}
