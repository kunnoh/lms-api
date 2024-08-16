package response

type UserResponse struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	IdNumber string `json:"idNumber"`
}
