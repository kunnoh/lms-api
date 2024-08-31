package utils

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}) (string, error) {
	// Load the private key
	privateKey, err := loadPrivateKey("./keys/ecdsa_private_key.pem")

	if err != nil {
		return "", fmt.Errorf("error loading private key: %w", err)
	}

	now := time.Now().UTC()
	claims := jwt.MapClaims{
		"exp": now.Add(ttl).Unix(),
		"iat": now.Unix(),
		"sub": payload,
		// "name": payload,
		"admin": true,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	// generate signed token str
	tokenStr, err := token.SignedString(privateKey)

	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}
	return tokenStr, nil
}

func ValidateToken(tokenString string, publicKey *ecdsa.PublicKey) (*jwt.Token, error) {
	// Load the public key
	publicKey, err := loadPublicKey("./keys/ecdsa_public_key.pem")
	if err != nil {
		return nil, fmt.Errorf("error loading public key")
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
