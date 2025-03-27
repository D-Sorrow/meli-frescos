package config

import (
	"log"
	"os"
	"sync"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/joho/godotenv"
	"github.com/melisource/fury_go-toolkit-secrets/pkg/secrets"
)

type Config struct {
	User          string
	Passwd        string
	MysqlEndpoint string
	DBName        string
}

var (
	config *mysql.Config
	once   sync.Once
)

func NewConfig() (*mysql.Config, error) {
	var err error
	once.Do(func() {
		client, err := secrets.NewClient()
		if err != nil {
			log.Fatal(err)
			return
		}

		mysqlEndpoint := os.Getenv("DB_MYSQL_DESAENV10_BGOW15S436_BGOW15S436_ENDPOINT")
		dbName := "bgow15s436"

		dbUser, ok := client.GetSecret("DB_MYSQL_DESAENV10_BGOW15S436_BGOW15S436_WPROD_USER")
		if !ok {
			log.Fatal("[dbUser] Secret not found")
		}

		dbPassword, ok := client.GetSecret("DB_MYSQL_DESAENV10_BGOW15S436_BGOW15S436_WPROD")
		if !ok {
			log.Fatal("[dbPassword] Secret not found")
		}

		config = &mysql.Config{
			User:                 dbUser,
			Passwd:               dbPassword,
			Net:                  "tcp",
			Addr:                 mysqlEndpoint,
			DBName:               dbName,
			Timeout:              500 * time.Millisecond,
			ReadTimeout:          500 * time.Millisecond,
			WriteTimeout:         500 * time.Millisecond,
			ParseTime:            true,
			AllowNativePasswords: true,
		}

	})
	return config, err
}
