package response

type UserResponse struct {
	UserId   int    `json:"userId"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    int    `json:"phone"`
	IdNumber int    `json:"idNumber"`
}
