package services

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
	"github.com/kunnoh/lms-api/src/utils"
)

type AuthService interface {
	Login(user request.LoginRequest) response.Response
	Register(user request.CreateUserRequest) response.Response
}

type AuthServiceImpl struct {
	UserRepo repository.UserRepository
	Validate *validator.Validate
}

func NewAuthServiceImpl(userRepo repository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepo: userRepo,
		Validate: validate,
	}
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(user request.LoginRequest) response.Response {
	err := a.Validate.Struct(user)
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation failed",
			Error:  err.Error(),
		}
	}

	u, err := a.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Error logging user",
			Error:  err.Error(),
		}
	}

	verify_err := utils.VerifyPassword(u.Password, user.Password)
	if verify_err != nil {
		return response.Response{
			Code:   http.StatusUnauthorized,
			Status: "Error logging user",
			Error:  verify_err.Error(),
		}
	}

	config, _ := config.LoadConfig(".")

	// generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, u.UserId, config.TokenSecret)

	if err_token != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "authorization failed",
			Error:  err_token.Error(),
		}
	} else {
		res := response.LoginResponse{
			TokenType: "Bearer",
			Token:     token,
		}
		return response.Response{
			Code:   http.StatusOK,
			Status: "authorized",
			Data:   res,
		}
	}
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(user request.CreateUserRequest) response.Response {
	err := a.Validate.Struct(user)
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation failed",
			Error:  err.Error(),
		}
	}

	hashedPW, err := utils.HashPassword(user.Password)
	utils.ErrorPanic(err)

	newUser := model.User{
		Email:    user.Email,
		Password: hashedPW,
		Name:     user.Name,
		Phone:    user.Phone,
		IdNumber: user.IdNumber,
	}

	errr := a.UserRepo.Save(newUser)
	if errr != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error saving user",
			Error:  err.Error(),
		}
	}
	return response.Response{
		Code:   http.StatusCreated,
		Status: "success",
	}
}
