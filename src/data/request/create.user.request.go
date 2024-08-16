package request

type CreateUserRequest struct {
	Name  string `validate:"required,min=1,max=10" json:"name"`
	Email string `validate:"required,min=6,max=20" json:"email"`
}
