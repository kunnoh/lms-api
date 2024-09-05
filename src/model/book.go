package model

import "github.com/google/uuid"

type Book struct {
	BookId      uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Title       string    `gorm:"type:varchar(255);"`
	ISBN        string    `gorm:"type:varchar(255);"`
	Publication string    `gorm:"type:varchar(255);"`
	Genre       string    `gorm:"type:varchar(255);"`
	Author      string    `gorm:"type:varchar(255);"`
}

func BookModel() {

}
