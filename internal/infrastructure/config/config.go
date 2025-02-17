package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	_ "github.com/joho/godotenv"
)

type Config struct {
	User   string
	Passwd string
	Net    string
	Addr   string
	Port   int
	DBName string
}

var (
	config *Config
	once   sync.Once
)

func NewConfig() (*Config, error) {
	var err error
	once.Do(func() {
		var (
			host    = os.Getenv("DB_HOST")
			port    = os.Getenv("DB_PORT")
			user    = os.Getenv("DB_USER")
			pwd     = os.Getenv("DB_PASSWORD")
			db_name = os.Getenv("DB_NAME")
		)

		portInt, err := strconv.Atoi(port)

		if err != nil {
			log.Fatal(err)
			return
		}

		config = &Config{
			User:   user,
			Passwd: pwd,
			Net:    "tcp",
			Addr:   host,
			Port:   portInt,
			DBName: db_name,
		}

	})
	return config, err
}
