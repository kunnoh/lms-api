package services

import (
	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/data/response"
)

type UserService interface {
	Create(user request.CreateUserRequest) response.Response
	Update(user request.UpdateUserRequest) response.Response
	Delete(UserId string) response.Response
	FindById(UserId string) response.Response
	FindAll() response.Response
}
