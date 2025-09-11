package server

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DBConn *sql.DB

func Init(dbConnectionString string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
		return nil, err
	}
	return db, nil
}
