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

	// load environment vars
	confg, c_err := config.LoadConfig(".")
	if c_err != nil {
		utils.ErrorPanic(c_err)
	}

	// Connect DB
	db := config.DbConnection(&confg)
	db.Table("users").AutoMigrate(&model.User{})

	// repository
	userRepo := repository.NewUserServiceImpl(db)

	// service
	validate := validator.New()
	userService := services.NewUserServiceImpl(userRepo, validate)
	authService := services.NewAuthServiceImpl(userRepo, validate)

	// controller
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)

	// routes
	route := routes.NewRouter(userController, authController)

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
