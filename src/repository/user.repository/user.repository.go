package userrepository

import "github.com/kunnoh/lms-api/src/model"

type UserRepository interface {
	Save(user model.User) (model.User, error)
	Update(user model.User)
	Delete(userId string) error
	FindById(userId string) (model.User, error)
	FindByEmail(email string) (model.User, error)
	FindAll() ([]model.User, error)
}
