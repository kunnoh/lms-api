package userrepository

import "github.com/kunnoh/lms-api/src/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	Update(user model.User)
	Delete(userId string) error
	FindById(userId string) (user model.User, err error)
	FindByEmail(email string) (model.User, error)
	FindAll() (users []model.User, err error)
}