package main

import (
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
	routes "github.com/kunnoh/lms-api/src/router"
	"github.com/kunnoh/lms-api/src/services"
	"github.com/kunnoh/lms-api/src/utils"
)

func main() {
	log.Info().Msg("started server")

	// Connect DB
	db := config.DbConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.User{})

	// repository
	userRepo := repository.NewUserServiceImpl(db)

	// service
	userService := services.NewUserServiceImpl(userRepo, validate)

	// controller
	userController := controller.NewUserController(userService)

	// routes
	route := routes.NewRouter(userController)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7755"
	}
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: route,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
