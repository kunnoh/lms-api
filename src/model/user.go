package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserId   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	Email    string    `gorm:"type:varchar(100);unique;not null"`
	Name     string    `gorm:"type:varchar(100);not null"`
	Password string    `gorm:"type:varchar(255);not null"`
	Phone    string    `gorm:"type:varchar(20);unique;not null"`
	IdNumber string    `gorm:"type:varchar(20);unique;not null"`
}

// hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserId = uuid.New()
	return
}
