package services

import (
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

func NewUserServiceImpl(userRepository repository.UserRepository, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo: userRepository,
		validate: validate,
	}
}

// Create implements UserService.
func (u *UserServiceImpl) Create(user request.CreateUserRequest) {
	err := u.validate.Struct(user)
	utils.ErrorPanic(err)
	useModel := model.User{
		Name: user.Name,
	}
	u.UserRepo.Save(useModel)
}

// Delete implements UserService.
func (u *UserServiceImpl) Delete(UserId int) {
	u.UserRepo.Delete(UserId)
}

// FindAll implements UserService.
func (u *UserServiceImpl) FindAll() []response.UserResponse {
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
	return users
}

// FindById implements UserService.
func (u *UserServiceImpl) FindById(UserId int) {
	userData, err := u.UserRepo.FindById(UserId)
	utils.ErrorPanic(err)

	userRes := response.UserResponse{
		UserId:   userData.UserId,
		Name:     userData.Name,
		Email:    userData.Email,
		Phone:    userData.Phone,
		IdNumber: userData.IdNumber,
	}
	return userRes
}

// Update implements UserService.
func (u *UserServiceImpl) Update(user request.UpdateUserRequest) {
	userData, err := u.UserRepo.FindById(user.UserId)
	utils.ErrorPanic(err)
	userData.Name = user.Name
	u.UserRepo.Update(userData)
}
