package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type CloudflareConfig struct {
	ApiKey  string `yaml:"api-key"`
	Records []struct {
		Schedule string `yaml:"schedule"`
		Name     string `yaml:"name"`
		Id       string `yaml:"id"`
		ZoneID   string `yaml:"zone-id"`
	}
}

type Config struct {
	Cloudflare CloudflareConfig
	IpProvider string `yaml:"ip-provider"`
}

func getConfig(configPath string) *Config {
	config := Config{}
	// fmt.Println("Using config file: " + configPath)

	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Error reading config file, please create/supply a config file")
		panic(err)
	}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		panic(err)
	}

	return &config
}