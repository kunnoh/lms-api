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
}

// LoadConfig retrieves configuration from environment variables
func LoadConfig() (config Config, err error) {
	config.DBHost = os.Getenv("DB_HOST")
	// config.DBPort = os.Getenv("DB_PORT")
	config.DBUsername = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")
	// config.Port = os.Getenv("PORT")
	config.TokenAge = os.Getenv("TOKEN_MAXAGE")
	config.TokenSecret = os.Getenv("TOKEN_SECRET")

	// Get the application port from the environment variable
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return config, fmt.Errorf("unable to parse PORT: %w", err)
	}
	config.DBPort = dbport

	// Get the application port from the environment variable
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return config, fmt.Errorf("unable to parse PORT: %w", err)
	}
	config.Port = port

	// Get the refresh token expiry from the environment variable
	refreshTokenExpiry, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRY"))
	if err != nil {
		return config, fmt.Errorf("unable to parse TOKEN_EXPIRY: %w", err)
	}
	config.RefreshTokenExpiresIn = refreshTokenExpiry

	// Get the token expiry from the environment variable
	tokenExpiry, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRY"))
	if err != nil {
		return config, fmt.Errorf("unable to parse TOKEN_EXPIRY: %w", err)
	}
	config.TokenExpiresIn = tokenExpiry

	return config, nil
}
