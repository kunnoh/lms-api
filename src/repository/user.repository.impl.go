package repository

import (
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
func (u *UserServiceImpl) Delete(userId string) {
	var user model.User
	res := u.Db.Where("id = ?", userId).Delete(&user)
	utils.ErrorPanic(res.Error)

}

// FindAll implements UserRepository.
func (u *UserServiceImpl) FindAll() ([]model.User, error) {
	var users []model.User
	res := u.Db.Find(&users)

	// Handle any error that occurs during the database query
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

// FindById implements UserRepository.
func (u *UserServiceImpl) FindById(userId string) (model.User, error) {
	var userr model.User
	res := u.Db.First(&userr, userId)
	if res.Error != nil {
		return userr, res.Error
	}
	return userr, nil
}

// Save implements UserRepository.
func (u *UserServiceImpl) Save(user model.User) error {
	res := u.Db.Create(&user)
	fmt.Println(res)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// Update implements UserRepository.
func (u *UserServiceImpl) Update(user model.User) {
	var updateUser = request.UpdateUserRequest{
		// UserId: user.UserId,
		Name: user.Name,
	}
	res := u.Db.Model(&user).Updates(updateUser)
	utils.ErrorPanic(res.Error)
}

func NewUserServiceImpl(Db *gorm.DB) UserRepository {
	return &UserServiceImpl{Db: Db}
}
