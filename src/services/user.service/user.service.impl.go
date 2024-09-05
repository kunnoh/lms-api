package userservice

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/repository"
)

type UserServiceImpl struct {
	UserRepo repository.UserRepository
	validate *validator.Validate
}

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepository,
		validate: validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) response.Response {
	err := u.validate.Struct(user)
	if err != nil {
		return response.Response{
			Code:   http.StatusBadRequest,
			Status: "validation failed",
			Error:  err.Error(),
		}
	}

	newUser := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Phone:    user.Phone,
		IdNumber: user.IdNumber,
		Password: user.Password,
	}

	savedUser, err := u.UserRepo.Save(newUser)
	if err != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "Error saving user",
			Error:  err.Error(),
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

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId string) response.Response {
	res := u.UserRepo.Delete(UserId)
	if res != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "error",
			Data:   "Failed to delete user",
		}
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "user deleted",
	}
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() response.Response {
	res, err := u.UserRepo.FindAll()
	if err != nil {
		return response.Response{
			Code:   http.StatusInternalServerError,
			Status: "error",
			Data:   "Failed to fetch users",
		}
	}

	users := make([]response.UserResponse, 0, len(res))

	for _, val := range res {
		users = append(users, response.UserResponse{
			UserId:   val.UserId,
			Name:     val.Name,
			Email:    val.Email,
			IdNumber: val.IdNumber,
			Phone:    val.Phone,
		})
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data:   users,
	}
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(UserId string) response.Response {
	userData, err := u.UserRepo.FindById(UserId)
	if err != nil {
		return response.Response{
			Code:   http.StatusNotFound,
			Status: "Error",
			Error:  "User not found",
		}
	}

	return response.Response{
		Code:   http.StatusOK,
		Status: "success",
		Data: response.UserResponse{
			UserId:   userData.UserId,
			Name:     userData.Name,
			Email:    userData.Email,
			Phone:    userData.Phone,
			IdNumber: userData.IdNumber,
		},
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
