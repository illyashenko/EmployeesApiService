package appconfig

import (
	"encoding/json"
	"log"
	"os"
)

type RedisConfig struct {
	Address   string `json:"address"`
	Password  string `json:"password"`
	DefaultDB int    `json:"defaultDb"`
}

type AppConfig struct {
	RedisConfig RedisConfig `json:"redis"`
	KafkaMq     KafkaConfig `json:"kafka"`
}

type KafkaConfig struct {
	Config             []map[string]string `json:"Config"`
	AdditionalSettings string              `json:"AdditionalSettings"`
}

func (config *AppConfig) configure() {

	data, err := os.ReadFile("config.json")

	if err != nil {
		log.Fatal("Config file reading error", err)
	}

	err = json.Unmarshal(data, config)

	if err != nil {
		log.Fatal("Config json deserialization error", err)
	}
}

func NewConfig() AppConfig {

	appConfig := AppConfig{}
	appConfig.configure()

	return appConfig
}
