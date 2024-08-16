package services

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
	"github.com/kunnoh/lms-api/src/utils"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	validate *validator.Validate
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) response.Response {
	err := u.validate.Struct(user)
	utils.ErrorPanic(err)
	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		IdNumber: user.IdNumber,
		Password: user.Password,
	}
	fmt.Println(userModel)
	u.UserRepo.Save(userModel)

	return response.Response{
		Code:   http.StatusCreated,
		Status: "success",
	}
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId int) {
	u.UserRepo.Delete(UserId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() response.Response {
	res := u.UserRepo.FindAll()
	var users []response.UserResponse
	for _, val := range res {
		user := response.UserResponse{
			UserId:   val.UserId,
			Name:     val.Name,
			Email:    val.Email,
			IdNumber: val.IdNumber,
			Phone:    val.Phone,
		}
		users = append(users, user)
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   users,
	}
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(UserId int) response.Response {
	userData, err := u.UserRepo.FindById(UserId)
	if err != nil {
		return response.Response{
			Code:   http.StatusNotFound,
			Status: "error",
			Error:  "User not found",
		}
	}

	userRes := response.UserResponse{
		UserId:   userData.UserId,
		Name:     userData.Name,
		Email:    userData.Email,
		Phone:    userData.Phone,
		IdNumber: userData.IdNumber,
	}
	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   userRes,
	}
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) response.Response {
	err := u.validate.Struct(user)
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation error",
			Error:  err.Error(),
		}
	}

	userData, err := u.UserRepo.FindById(user.UserId)
	if err != nil {
		return response.Response{
			Code:   http.StatusNotFound,
			Status: "error",
			Error:  "User not found",
		}
	}

	userData.Name = user.Name
	// Update other fields as necessary
	u.UserRepo.Update(userData)
	return response.Response{
		Code:   http.StatusOK,
		Status: "User updated successfully",
	}
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepository,
		validate: validate,
	}
}
