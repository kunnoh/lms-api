package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(c *Config) *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	}

	// verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Could not verify database connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}
