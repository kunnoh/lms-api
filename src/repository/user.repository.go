package repository

import "github.com/kunnoh/lms-api/src/model"

type UserRepository interface {
	Save(user model.User)
	Update(user model.User)
	Delete(userId int)
	FindById(userId int) (user model.User, err error)
	FindAll() []model.User
}
