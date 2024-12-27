package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// custom claims
type CustomClaims struct {
	Role string `json:"role"`
	jwt.Claims
}

func GenerateToken(ttl time.Duration, userId uuid.UUID) (string, error) {
	// Load the private key
	PrivKey := os.Getenv("PRIV_KEY")
	// privateKey, err := loadPrivateKey("./keys/ecdsa_private_key.pem")

	// if err != nil {
	// 	return "", fmt.Errorf("error loading private key: %w", err)
	// }

	token := jwt.New(jwt.SigningMethodES256)
	now := time.Now().UTC()

	claims := token.Claims.(jwt.MapClaims)
	iat := now.Unix()
	exp := now.Add(ttl).Unix()

	claims["role"] = "moderator"
	claims["exp"] = exp
	claims["iat"] = iat
	claims["iss"] = "https://lmsapi.com"
	claims["aud"] = "lms-api"
	claims["sub"] = userId

	// generate signed token str
	tokenStr, err := token.SignedString(PrivKey)

	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	// Load the public key
	PubKey := os.Getenv("PUB_KEY")
	// publicKey, err := loadPublicKey("./keys/ecdsa_public_key.pem")
	// if err != nil {
	// 	return nil, fmt.Errorf("error loading public key: %v", err)
	// }

	// Verify token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return PubKey, nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("token is expired")
	}

	return token, nil
}
