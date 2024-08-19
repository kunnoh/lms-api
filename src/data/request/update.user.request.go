package request

type UpdateUserRequest struct {
	UserId string `valildate:"required"`
	Name   string `validate:"required,max=200,min=1" json:"name"`
	Email  string `validate:"required,max=50,min=5" json:"email"`
	Phone  string `validate:"required,max=15,min=10" json:"phone"`
}
