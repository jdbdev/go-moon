package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Config holds database configuration
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// Connect establishes a connection to the PostgreSQL database
func Connect(cfg Config) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	// Test the connection
	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %w", err)
	}

	log.Println("Successfully connected to database")
	return nil
}

// Close closes the database connection
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
