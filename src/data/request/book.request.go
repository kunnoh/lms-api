package request

type CreateBookRequest struct {
	Title       string `validate:"required,min=3,max=40" json:"title"`
	ISBN        string `validate:"required,min=6,max=20" json:"isbn"`
	Publication string `validate:"required,min=8,max=15" json:"publication"`
	Genre       string `validate:"required,min=8,max=30" json:"genre"`
	Author      string `validate:"required,min=8,max=40" json:"author"`
}

// type UpdateBookRequest struct {
// 	Title       string `validate:"required,min=3,max=40" json:"title"`
// 	ISBN        string `validate:"required,min=6,max=20" json:"isbn"`
// 	Publication string `validate:"required,min=8,max=15" json:"publication"`
// 	Genre       string `validate:"required,min=8,max=30" json:"genre"`
// 	Author      string `validate:"required,min=8,max=40" json:"author"`
// 	// CreatedAt   time.Duration `validate:"required" json:"author`
// }
