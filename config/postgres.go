package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

// WritePostgreDB - function for creating database connection for write-access
func WritePostgresDB() *sqlx.DB { 
	return CreatePostgresDBConnection(fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
		os.Getenv("WRITE_DB_USER"), os.Getenv("WRITE_DB_NAME"), os.Getenv("WRITE_DB_PASSWORD"), os.Getenv("WRITE_DB_HOST")))
}

// ReadPostgreDB function for creating database connection for read-access
func ReadPostgresDB() *sqlx.DB {
  return CreatePostgresDBConnection(fmt.Sprintf("user=%s dbname=%s sslmode=disable password=%s host=%s",
  os.Getenv("READ_DB_USER"), os.Getenv("READ_DB_NAME"), os.Getenv("READ_DB_PASSWORD"), os.Getenv("READ_DB_HOST")))
}

// CreateDBConnection function for creating database connection
func CreatePostgresDBConnection(descriptor string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", descriptor)
	if err != nil {
		log.Fatalf("error connecting to DB: %s", err)
	}
	return db
}
