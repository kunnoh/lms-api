package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(ttl time.Duration, payload interface{}) (string, error) {
	privateKey, err := loadKey("./keys/ecdsa_private_key.pem")

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

func ValidateToken(token string, signedKey string) (interface{}, error) {
	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return []byte(signedKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)

	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}
	return claims["sub"], nil
}
