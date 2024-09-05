package request

type CreateBookRequest struct {
	Title    string `validate:"required,min=3,max=20" json:"title"`
	Email    string `validate:"required,min=6,max=20" json:"email"`
	Phone    string `validate:"required,min=8,max=15" json:"phone"`
	IdNumber string `validate:"required,min=8,max=12" json:"idnumber"`
}

type UpdateBookRequest struct {
	Title string `validate:"required,min=3,max=20" json:"title"`
}
