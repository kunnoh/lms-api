package services

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/config"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
	"github.com/kunnoh/lms-api/src/utils"
)

type AuthService interface {
	Login(user request.LoginRequest) (string, error)
	Register(user request.CreateUserRequest)
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
func (a *AuthServiceImpl) Login(user request.LoginRequest) (string, error) {
	u, err := a.UserRepo.FindByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid username or password")
	}
	config, _ := config.LoadConfig(".")

	verify_err := utils.VerifyPassword(u.Password, user.Password)
	if verify_err != nil {
		return "", errors.New("invalid username or password")
	}

	// generate token
	token, err_token := utils.GenerateToken(config.TokenExpiresIn, u.UserId, config.TokenSecret)
	utils.ErrorPanic(err_token)
	return token, nil
}

// Register implements AuthService.
func (a *AuthServiceImpl) Register(user request.CreateUserRequest) {
	hashedPW, err := utils.HashPassword(user.Password)
	utils.ErrorPanic(err)

	newUser := model.User{
		Email:    user.Email,
		Password: hashedPW,
		Name:     user.Name,
		Phone:    user.Phone,
		IdNumber: user.IdNumber,
	}

	a.UserRepo.Save(newUser)
}
