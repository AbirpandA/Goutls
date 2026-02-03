package godbutils

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Postgres driver
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func ConnectDb(cfg Config) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)
	var db *sql.DB
	var err error

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to the database!")
	return db, nil
}
