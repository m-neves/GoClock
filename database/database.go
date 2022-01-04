package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	connStr = `
		user=%s
		password=%s
		host=%s
		port=%s
		dbname=%s
		search_path=%s
		connect_timeout=%s
		sslmode=%s		
	`
)

var db *sql.DB = nil

func Connect() *sql.DB {
	db, err := sql.Open("postgres",
		fmt.Sprintf(
			connStr,
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SCHEMA"),
			os.Getenv("DB_CONNECT_TIMEOUT"),
			os.Getenv("DB_SSL_MODE"),
		))

	if err != nil {
		log.Fatalf("Could not initialize database with message: %s", err.Error())
	}

	return db
}

func GetDb() *sql.DB {
	if db == nil {
		db = Connect()
	}

	return db
}
