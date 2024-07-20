package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"

	"github.com/kunnoh/lms-api/src/utils"
)

func main() {
	log.Info().Msg("started server")
	routes := gin.Default()

	routes.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Welcome to LMS-API")
	})

	server := &http.Server{
		Addr:    ":7755",
		Handler: routes,
	}

	err := server.ListenAndServe()
	fmt.Println(err)
	utils.ErrorPanic(err)
}
