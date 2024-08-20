package repository

import "github.com/kunnoh/lms-api/src/model"

type UserRepository interface {
	Save(user model.User) (err error)
	Update(user model.User)
	Delete(userId string)
	FindById(userId string) (user model.User, err error)
	FindByEmail(email string) (model.User, error)
	FindAll() (users []model.User, err error)
}
