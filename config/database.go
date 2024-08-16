package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/kunnoh/lms-api/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {
	// Load environment variables from the .env file
	err := godotenv.Load()
	utils.ErrorPanic(err)

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	fmt.Println(host)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
