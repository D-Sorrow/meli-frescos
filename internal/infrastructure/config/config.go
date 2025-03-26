package config

import (
	"log"
	"os"
	"strconv"
	"strings"
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
	TLS    bool
}

var (
	config *Config
	once   sync.Once
)

func stringToBool(s string) bool {
	s = strings.ToLower(s)

	switch s {
	case "1", "on", "true", "t", "yes", "y":
		return true
	case "0", "off", "false", "f", "no", "n":
		return false
	default:
		return true
	}
}

func NewConfig() (*Config, error) {
	var err error
	once.Do(func() {
		var (
			host   = os.Getenv("DB_HOST")
			port   = os.Getenv("DB_PORT")
			user   = os.Getenv("DB_USER")
			pwd    = os.Getenv("DB_PASSWORD")
			dbName = os.Getenv("DB_NAME")
			tlsStr = os.Getenv("DB_TLS")
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
			DBName: dbName,
			TLS:    stringToBool(tlsStr),
		}

	})
	return config, err
}
