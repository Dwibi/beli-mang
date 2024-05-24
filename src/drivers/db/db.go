package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func CreateConnection() (*sql.DB, error) {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbParams := os.Getenv("DB_PARAMS")

	strConnection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?%s", dbUsername, dbPassword, dbHost, dbPort, dbName, dbParams)

	// Define connection pool parameters (adjust as needed)
	maxOpenConns := 25 // Maximum number of open connections in the pool
	maxIdleConns := 25 // Maximum number of idle connections in the pool

	db, err := sql.Open("postgres", strConnection)
	if err != nil {
		return nil, err
	}

	// Create connection pool
	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	// Test connection using PingContext
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		db.Close() // Close the connection pool on error
		return nil, err
	}

	log.Println("Database connected!")

	return db, nil
}
