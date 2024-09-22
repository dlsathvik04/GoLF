package config

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Config struct {
	Servers             []string `json:"servers"`
	Capacities          []int    `json:"capacities"`
	Algorithm           int      `json:"lbAlgorithm"`
	Port                string   `json:"port"`
	HealthCheckInterval string   `json:"healthCheckInterval"`
}

func GetConfig() *Config {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	var config Config
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}
	return &config
}
