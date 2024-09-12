package main

import (
	logg "log"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/model"
	bookrepository "github.com/kunnoh/lms-api/src/repository/book.repository"
	userrepository "github.com/kunnoh/lms-api/src/repository/user.repository"
	routes "github.com/kunnoh/lms-api/src/router"
	"github.com/kunnoh/lms-api/src/services"
	bookservice "github.com/kunnoh/lms-api/src/services/book.service"
	userservice "github.com/kunnoh/lms-api/src/services/user.service"
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
	db, db_err := config.DbConnection(&confg)
	if db_err != nil {
		logg.Fatalf("ERROR: %v", db_err)
	}

	if err := db.Table("users").AutoMigrate(&model.User{}); err != nil {
		utils.ErrorPanic(err)
	}

	if err := db.Table("books").AutoMigrate(&model.Book{}); err != nil {
		utils.ErrorPanic(err)
	}

	// repository
	userRepo := userrepository.NewUserServiceImpl(db)
	bookRepo := bookrepository.NewBookRepositoryImpl(db)

	// service
	validate := validator.New()
	userService := userservice.NewUserServiceImpl(userRepo, validate)
	bookService := bookservice.NewBookServiceImpl(bookRepo, validate)
	authService := services.NewAuthServiceImpl(userRepo, validate)

	// controller
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	bookController := controller.NewBookController(bookService)

	// routes
	route := routes.NewRouter(userRepo, userController, bookController, authController, db)

	PORT := confg.Port

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(PORT),
		Handler: route,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
