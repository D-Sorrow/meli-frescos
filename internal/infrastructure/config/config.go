package config

import (
	"github.com/subosito/gotenv"
	_ "github.com/subosito/gotenv"
	"os"
	"sync"
)

type Config struct {
	ServerAddr string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var (
	config *Config
	once   sync.Once
)

func NewConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if err = gotenv.Load(); err != nil {

		}
		config = &Config{
			ServerAddr: os.Getenv("SERVER_ADDR"),
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		}

	})
	return config, err
}
