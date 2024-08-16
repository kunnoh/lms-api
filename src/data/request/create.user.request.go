package request

type CreateUserRequest struct {
	Name     string `validate:"required,min=1,max=10" json:"name"`
	Email    string `validate:"required,min=3,max=20" json:"email"`
	Password string `validate:"required,min=6,max=40" json:"password"`
	Phone    string `validate:"required,min=8,max=15" json:"phone"`
	IdNumber string `validate:"required,min=8,max=12" json:"idnumber"`
}
