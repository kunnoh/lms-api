package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost                string
	DBPort                int
	DBUsername            string
	DBPassword            string
	DBName                string
	Port                  int
	RefreshTokenExpiresIn time.Duration
	TokenExpiresIn        time.Duration
	TokenAge              string
	TokenSecret           string
	PrivKey               string
	PubKey                string
}

// LoadConfig retrieves configuration from environment variables
func LoadConfig() (config Config, err error) {
	config.DBHost = os.Getenv("DB_HOST")
	config.DBUsername = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	config.TokenAge = os.Getenv("TOKEN_MAXAGE")
	config.PrivKey = os.Getenv("PRIV_KEY")
	config.PubKey = os.Getenv("PUB_KEY")

	// Check if all required environment variables are set

	// if config.DBHost == "" || config.DBUsername == "" || config.DBPassword == "" || config.DBName == "" || config.TokenSecret == "" {
	// 	return config, fmt.Errorf("required environment variables are missing")
	// }

	// db port
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return config, fmt.Errorf("unable to parse DB_PORT: %w", err)
	}
	config.DBPort = dbport

	// app port
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return config, fmt.Errorf("unable to parse PORT: %w", err)
	}
	config.Port = port

	//refresh token expiry
	refreshTokenExpiry, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		return config, fmt.Errorf("unable to parse REFRESH_TOKEN_EXPIRY: %w", err)
	}
	config.RefreshTokenExpiresIn = refreshTokenExpiry

	// token expiry
	tokenExpiry, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRY"))
	if err != nil {
		return config, fmt.Errorf("unable to parse TOKEN_EXPIRY: %w", err)
	}
	config.TokenExpiresIn = tokenExpiry

	return config, nil
}
