package response

import (
	"github.com/google/uuid"
)

type UserResponse struct {
	UserId   uuid.UUID `json:"userId"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	IdNumber string    `json:"idNumber"`
}

type LoginResponse struct {
	TokenType string `json:"token_type"`
	Token     string `json:"token"`
}
