package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "moderator"
	Client    Role = "client"
)

type User struct {
	// UserId    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	UserId    uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Phone     string    `gorm:"type:varchar(20);unique;not null"`
	IdNumber  string    `gorm:"type:varchar(20);unique;not null"`
	Role      Role      `gorm:"type:varchar(20);default:'client'"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// hook to generate UUID
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.UserId = uuid.New()
	return
}
