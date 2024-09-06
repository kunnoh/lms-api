package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	// BookId uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey;"`
	BookId      uuid.UUID `gorm:"type:uuid;primary_key;"`
	Title       string    `gorm:"type:varchar(255);unique;not null"`
	ISBN        string    `gorm:"type:varchar(255);unique;not null"`
	Publication string    `gorm:"type:varchar(255);not null"`
	Genre       string    `gorm:"type:varchar(255);not null"`
	Author      string    `gorm:"type:varchar(255);not null"`
}

// hook to generate UUID
func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	b.BookId = uuid.New()
	return
}
