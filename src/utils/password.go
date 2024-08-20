package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash the password")
	}
	return string(hashPass), nil
}

func verifyPassword(hashPass string, pass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPass), []byte((pass)))
}
