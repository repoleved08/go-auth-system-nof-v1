package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	
	connString := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),os.Getenv("DB_NAME"),os.Getenv("DB_PORT"),os.Getenv("DB_HOST"),)

	DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Failed to connect to the db: %v", err)
	}

	log.Println("Connection to DB was successful!")
}
