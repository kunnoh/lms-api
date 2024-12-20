package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	authservice "github.com/kunnoh/lms-api/src/services/auth.service"
)

type AuthController struct {
	authService authservice.AuthService
}

func NewAuthController(service authservice.AuthService) *AuthController {
	return &AuthController{authService: service}
}

func (ctrl *AuthController) Login(ctx *gin.Context) {
	var loginReq request.LoginRequest

	if err := ctx.ShouldBindJSON(&loginReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad login request",
			Error:  err.Error(),
		})
		return
	}

	res := ctrl.authService.Login(loginReq)
	ctx.JSON(res.Code, res)
}

func (ctrl *AuthController) Register(ctx *gin.Context) {
	var registerReq request.CreateUserRequest

	if err := ctx.ShouldBindJSON(&registerReq); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Response{
			Code:   http.StatusBadRequest,
			Status: "Bad registration request",
			Error:  err.Error(),
		})
		return
	}

	res := ctrl.authService.Register(registerReq)
	ctx.JSON(res.Code, res)
}

func (ctrl *AuthController) RefreshToken(ctx *gin.Context) {

}
