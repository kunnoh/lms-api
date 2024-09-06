package services

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/model"
	userrepository "github.com/kunnoh/lms-api/src/repository/user.repository"
	"github.com/kunnoh/lms-api/src/utils"
)

type AuthService interface {
	Login(user request.LoginRequest) response.Response
	Register(user request.CreateUserRequest) response.Response
	RefreshToken(user request.CreateUserRequest) response.Response
}

type AuthServiceImpl struct {
	UserRepo userrepository.UserRepository
	Validate *validator.Validate
}

func NewAuthServiceImpl(userRepo userrepository.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepo: userRepo,
		Validate: validate,
		// Config: config.Config,
	}
}

// Login implements AuthService.
func (a *AuthServiceImpl) Login(user request.LoginRequest) response.Response {
	err := a.Validate.Struct(user)
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "fail",
			Error:  err.Error(),
		}
	}

	u, err := a.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return response.Response{
			Code:   http.StatusUnauthorized,
			Status: "fail",
			Error:  "Incorrect username or password",
		}
	}

	verify_err := utils.VerifyPassword(u.Password, user.Password)
	if verify_err != nil {
		return response.Response{
			Code:   http.StatusUnauthorized,
			Status: "fail",
			Error:  verify_err.Error(),
		}
	}

	config, _ := config.LoadConfig(".")

	// generate token
	token, err_tok := utils.GenerateToken(config.TokenExpiresIn, u.UserId)
	refreshtoken, err_token := utils.GenerateToken(config.RefreshTokenExpiresIn, u.UserId)

	if err_token != nil || err_tok != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation failed",
			Error:  err_token.Error(),
		}
	} else {
		res := response.LoginResponse{
			TokenType:    "Bearer",
			Token:        token,
			RefreshToken: refreshtoken,
		}
		return response.Response{
			Code:   http.StatusOK,
			Status: "success",
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
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "Hashing password failed",
			Error:  err.Error(),
		}
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		IdNumber: user.IdNumber,
		Password: string(hashedPW),
	}

	savedUser, errr := a.UserRepo.Save(newUser)
	if errr != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error saving user",
			Error:  errr.Error(),
		}
	}

	return response.Response{
		Code:   http.StatusCreated,
		Status: "success",
		Data: response.UserResponse{
			UserId:   savedUser.UserId,
			Email:    savedUser.Email,
			Name:     savedUser.Name,
			Phone:    savedUser.Phone,
			IdNumber: savedUser.IdNumber,
		},
	}
}

// Refresh impements AuthService.
func (a *AuthServiceImpl) RefreshToken(user request.CreateUserRequest) response.Response {
	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
	}
}
