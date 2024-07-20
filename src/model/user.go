package model

type User struct {
	Id         int    `gorm:"type:int;primary_key"`
	Email      string `gorm:"type:varchar(255);uniq_key"`
	Password   string `gorm:"type:varchar(255);"`
	Phone      int    `gorm:"type:int;primary_key"`
	SchoolCode string `gorm:"type:int;primary_key"`
	Profession string `gorm:"type:varchar(255);"`
}

func UserModel() {

}
