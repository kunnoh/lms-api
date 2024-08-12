package services

import (
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest)
	Update(user request.UpdateUserRequest)
	Delete(UserId int)
	FindById(UserId int) response.UserResponse
	FindAll() []response.UserResponse
}
