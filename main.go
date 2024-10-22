package main

import (
	"context"
	logg "log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	log.Info().Msg("Started server")

	// load environment vars
	confg, c_err := config.LoadConfig()
	if c_err != nil {
		utils.ErrorPanic(c_err)
	}

	PORT := confg.Port
	log.Info().Msgf("Server running on port: %v", PORT)

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

	server := &http.Server{
		Addr:         ":" + PORT,
		Handler:      route,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// listen for os signal
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt, syscall.SIGTERM)

	// Run server in a goroutine so it doesn't block
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logg.Fatalf("Could not listen on port %v: %v\n", PORT, err)
		}
	}()

	// block until a signal
	<-stopCh
	log.Info().Msg("Shutting down server...")

	// Create a deadline for the server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Gracefully shutdown the server
	if err := server.Shutdown(ctx); err != nil {
		logg.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info().Msg("Server exiting!!")
}
