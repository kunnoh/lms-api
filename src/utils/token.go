package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

// custom claims
type CustomClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(ttl time.Duration, userId uuid.UUID) (string, error) {
	// Load the private key
	privateKey, err := loadPrivateKey("./keys/ecdsa_private_key.pem")

	if err != nil {
		return "", fmt.Errorf("error loading private key: %w", err)
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"exp":  now.Add(ttl).Unix(),
		"nbf":  now.Unix(),
		"iat":  now.Unix(),
		"iss":  "https://lmsapi.com",
		"aud":  "lms-api",
		"sub":  userId,
		"role": "moderator",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// generate signed token str
	tokenStr, err := token.SignedString(privateKey)

	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	// Load the public key
	publicKey, err := loadPublicKey("./keys/ecdsa_public_key.pem")
	if err != nil {
		return nil, fmt.Errorf("error loading public key: %v", err)
	}

	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
