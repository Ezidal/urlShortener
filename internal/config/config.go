package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Environment  string `yaml:"environment" env-required:"true"`
	Storage_path string `yaml:"storage_path"`
	HttpServer   `yaml:"http_server"`
}

type HttpServer struct {
	Host         string        `yaml:"host"`
	Port         string        `yaml:"port"`
	Timeout      time.Duration `yaml:"timeout"`
	Idle_timeout time.Duration `yaml:"idle_timeout"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		// configPath = "./config/local.yaml"
		log.Fatal("CONFIG_PATH is not set")
	}

	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
