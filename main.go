package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/kunnoh/lms-api/src/utils"
)

func main() {
	log.Info().Msg("started server")

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
	fmt.Println(err)
	utils.ErrorPanic(err)
}
