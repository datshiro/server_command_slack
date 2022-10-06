package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

func NewConfig() Config {
	config := Config{}
	err := envconfig.Process("server", &config)
	if err != nil {
		log.Fatal(err.Error())
	}
	return config
}

type Config struct {
	Port    string `envconfig:"SERVER_PORT" default:"8080"`
	Host    string `envconfig:"SERVER_HOST" default:"localhost"`
	APIPath string `envconfig:"API_PATH" default:"/api/v1"`
}

func (c Config) ServerAddress() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
