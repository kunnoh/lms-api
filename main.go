package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"

	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/controller"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
	"github.com/kunnoh/lms-api/src/services"
	"github.com/kunnoh/lms-api/src/utils"
)

func main() {
	log.Info().Msg("started server")

	// Connect DB
	db := config.DbConnection()
	validate := validator.New()

	db.Table("user").AutoMigrate(&model.User{})

	// repository
	userRepo := repository.NewUserServiceImpl(db)

	// service
	userService := services.UserServiceImpl(userRepo, validate)

	// controller
	userController := controller.NewUserController()

	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Message string }{Message: "Welcome to LMS-API"})
	})

	routes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7755"
	}
	server := &http.Server{
		Addr:    ":" + PORT,
		Handler: routes,
	}

	err := server.ListenAndServe()
	utils.ErrorPanic(err)
}
