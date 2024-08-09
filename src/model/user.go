package model

type User struct {
	UserId   int    `gorm:"type:int;primary_key"`
	Email    string `gorm:"type:varchar(255);uniq_key"`
	Name     string `gorm:"type:varchar(255);"`
	Password string `gorm:"type:varchar(255);"`
	Phone    int    `gorm:"type:int;uniq_key"`
	IdNumber int    `gorm:"type:int;uniq_key"`
}

func UserModel() {

}
