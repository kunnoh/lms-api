package model

type Book struct {
	BookId      int    `gorm:"type:int;primary_key"`
	Title       string `gorm:"type:varchar(255);"`
	ISDN        string `gorm:"type:varchar(255);"`
	Publication string `gorm:"type:varchar(255);"`
	Genre       string `gorm:"type:varchar(255);"`
	ISBN        string `gorm:"type:varchar(255);"`
	Author      string `gorm:"type:varchar(255);"`
}

func BookModel() {

}
