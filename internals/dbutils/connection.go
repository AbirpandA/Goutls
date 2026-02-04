package godbutils

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq" // Import the driver anonymously
)


type Config struct{
	Host  string
	Port int
	User string
	Password string
	DBName string
}


func ConnectDb(cfg Config) (*sql.DB , error) {
	constr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host,cfg.Port,cfg.User,cfg.Password,cfg.DBName)

	var db *sql.DB
	var err error

	for  i := 0;  i < 5;  i++ {
		db,err= sql.Open("postgres",constr)
		if err==nil{
			err=db.Ping()
		}

		if err == nil{
			log.Println("Successfully connected to the database!")
            return db, nil
		}
		log.Printf("Failed to connect (attempt %d/5), retrying in 2s...", i+1)
		time.Sleep(2*time.Second)


	}

	return db , fmt.Errorf("could not connect to DB after 5 attempts: %v", err)
}