package db

import (
	"database/sql"
	"fmt"
	"github.com/D-Sorrow/meli-frescos/internal/infrastructure/config"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DataBase struct {
	db *sql.DB
}

func Connect(cfg *config.Config) *sql.DB {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
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

func (db *DataBase) NewDataBase(cfg *config.Config) *DataBase {
	if db.db == nil {
		db.db = Connect(&config.Config{})
	}
	panic("Database connection initialized")
}
