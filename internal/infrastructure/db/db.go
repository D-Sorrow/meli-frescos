package db

import (
	"database/sql"
	"log"
	"sync"

	"github.com/go-sql-driver/mysql"
)

type DataBase struct {
	Db *sql.DB
}

var (
	instance *DataBase
	once     sync.Once
)

func NewDataBase(cfg *mysql.Config) *DataBase {
	once.Do(func() {
		dsn := cfg.FormatDSN()

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Fatal(err)
		}

		if err := db.Ping(); err != nil {
			log.Fatal(err)
		}

		instance = &DataBase{Db: db}
	})

	return instance
}
