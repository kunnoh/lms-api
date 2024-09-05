package userrepository

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

func NewUserServiceImpl(Db *gorm.DB) UserRepository {
	return &UserServiceImpl{Db: Db}
}

// Delete implements UserRepository.
func (u *UserServiceImpl) Delete(userId string) error {
	res := u.Db.Where("user_id = ?", userId).Delete(&model.User{})

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return fmt.Errorf("no user found with id %s", userId)
	}

	return nil
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
	var user model.User
	res := u.Db.First(&user, "user_id = ?", userId)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, res.Error
	}
	return user, nil
}

// FindByEmail implements UserRepository.
func (u *UserServiceImpl) FindByEmail(email string) (model.User, error) {
	var user model.User
	res := u.Db.First(&user, "email = ?", email)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return user, errors.New("user not found")
		}
		return user, res.Error
	}
	return user, nil
}

// Save implements UserRepository.
func (u *UserServiceImpl) Save(user model.User) (model.User, error) {
	res := u.Db.Create(&user)
	if res.Error != nil {
		return user, res.Error
	}
	return user, nil
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
