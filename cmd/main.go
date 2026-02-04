package main

import (
	"fmt"
	"log"

	"go-dbutils/internals/dbutils"
)

func main() {
	cfg := godbutils.Config{
		Host:     "localhost",
		Port:     5432,
		User:     "user",
		Password: "password",
		DBName:   "mydb",
	}

	conn, err := godbutils.ConnectDb(cfg)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer conn.Close()

	fmt.Println("Connection test successful!")
}
