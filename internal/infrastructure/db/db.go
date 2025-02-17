package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/config"
	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	Db *sql.DB
}

var (
	instance *DataBase
	once     sync.Once
)

func Connect(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&tls=true", cfg.User, cfg.Passwd, cfg.Addr, cfg.Port, cfg.DBName)

	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to DB")
	return db
}

func NewDataBase(cfg *config.Config) *DataBase {
	once.Do(func() {
		instance = &DataBase{Db: Connect(cfg)}
		instance.Db.SetMaxOpenConns(10)
		instance.Db.SetMaxIdleConns(5)
		instance.Db.SetConnMaxLifetime(0)
	})

	return instance
}
