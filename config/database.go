package config

import (
	"fmt"

	"github.com/kunnoh/lms-api/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DbConnection() *gorm.DB {
	dbUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	utils.ErrorPanic(err)

	return db
}
