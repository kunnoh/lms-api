package repository

import (
	"errors"
	"fmt"

	"github.com/kunnoh/lms-api/src/data/request"
	"github.com/kunnoh/lms-api/src/model"
	"github.com/kunnoh/lms-api/src/utils"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Db *gorm.DB
}

// Delete implements UserRepository.
func (u *UserServiceImpl) Delete(userId int) {
	var user model.User
	res := u.Db.Where("id = ?", userId).Delete(&user)
	utils.ErrorPanic(res.Error)

}

// FindAll implements UserRepository.
func (u *UserServiceImpl) FindAll() []model.User {
	var users []model.User
	var user model.User
	res := u.Db.Find(&user)
	fmt.Println(user)
	utils.ErrorPanic(res.Error)
	return users
}

// FindById implements UserRepository.
func (u *UserServiceImpl) FindById(userId int) (user model.User, err error) {
	var userr model.User
	res := u.Db.Find(&userr, userId)
	if res != nil {
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

// Save implements UserRepository.
func (u *UserServiceImpl) Save(user model.User) {
	res := u.Db.Create(user)
	utils.ErrorPanic(res.Error)
}

// Update implements UserRepository.
func (u *UserServiceImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		UserId: user.UserId,
		Name:   user.Name,
	}
	res := u.Db.Model(&user).Updates(updateUser)
	utils.ErrorPanic(res.Error)
}

func NewUserServiceImpl(Db *gorm.DB) UserRepository {
	return &UserServiceImpl{Db: Db}
}
