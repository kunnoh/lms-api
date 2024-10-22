package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(c *Config) (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)

	var db *gorm.DB
	var err error

	// Retry logic
	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
		if err == nil {
			break
		}
		waitTime := time.Duration(1<<i) * time.Second
		log.Printf("Attempt %d: Failed to connect to database, retrying in %v...", i+1, waitTime)
		time.Sleep(waitTime)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database after retries: %w", err)
	}

	// verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve database object: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	log.Println("Database connection established successfully")

	return db, nil
}
