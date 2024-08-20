package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/services"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(service services.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (controller *AuthController) Login(ctx *gin.Context) {

}

func (controller *AuthController) Register(ctx *gin.Context) {

}
