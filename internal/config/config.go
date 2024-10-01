package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"prod"`
	HTTPServer `yaml:"http_server"`
	Storage    `yaml:"storage"`
}

type HTTPServer struct {
	Address string `yaml:"address" env-default:"localhost:8080"`
}

type Storage struct {
	Url      string `yaml:"url" env-default:"localhost:8080"`
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"password" env-required:"true"`
}

// Func to load config from yaml
func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
