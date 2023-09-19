package config

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	nginxContainer string = "NGINX_CONTAINER"
	configDir      string = "CONFIG_DIR"
	kafkaBrokers   string = "KAFKA_BROKERS"
	kafkaTopic     string = "KAFKA_TOPIC"
)

type Configuration struct {
	ConfigDir      string
	NginxContainer string
	Brokers        []string
	Topic          string
}

var AppConfig Configuration

func Init() {
	kafkaBrokersList := getStringListFromEnv(kafkaBrokers, "kafka1:9092,kafka2:9093,kafka3:9094")
	AppConfig = Configuration{
		ConfigDir:      getEnvString(configDir, "/app/sites-enabled"),
		NginxContainer: getEnvString(nginxContainer, "gateway"),
		Brokers:        kafkaBrokersList,
		Topic:          getEnvString(kafkaTopic, "automation-events"),
	}
}

func getStringListFromEnv(envVarName, defaultValue string) []string {
	value := getEnvString(envVarName, defaultValue)
	return strings.Split(value, ",")
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, err := strconv.Atoi(value)
		if err == nil {
			return intVal
		}
	}
	log.Printf("Using default value for %s: %v", key, defaultValue)
	return defaultValue
}

func getEnvString(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	log.Printf("Using default value for %s: %s", key, defaultValue)
	return defaultValue
}
