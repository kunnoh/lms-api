package config

import (
	"fmt"

	"github.com/kunnoh/lms-api/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection(c *Config) *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", c.DBHost, c.DBPort, c.DBUsername, c.DBPassword, c.DBName)

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
