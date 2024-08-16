package model

type User struct {
	UserId   int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"type:varchar(255);unique;not null"`
	Name     string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Phone    string `gorm:"type:varchar(255);unique;not null"`
	IdNumber string `gorm:"type:varchar(255);unique;not null"`
}
