package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	DBHost                string        // Database host
	DBPort                string        // Database port
	DBUsername            string        // Database username
	DBPassword            string        // Database password
	DBName                string        // Database name
	Port                  int           // Application port
	RefreshTokenExpiresIn time.Duration // Refresh token expiry duration
	TokenExpiresIn        time.Duration // Token expiry duration
	TokenAge              int           // Maximum age of token
	TokenSecret           string        // Token secret
}

// LoadConfig retrieves configuration from environment variables
func LoadConfig() (config Config, err error) {
	config.DBHost = os.Getenv("DB_HOST")
	config.DBPort = os.Getenv("DB_PORT")
	config.DBUsername = os.Getenv("DB_USER")
	config.DBPassword = os.Getenv("DB_PASSWORD")
	config.DBName = os.Getenv("DB_NAME")

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

	// Get the maximum token age from the environment variable
	tokenAge, err := strconv.Atoi(os.Getenv("TOKEN_MAXAGE"))
	if err != nil {
		return config, fmt.Errorf("unable to parse TOKEN_MAXAGE: %w", err)
	}
	config.TokenAge = tokenAge

	config.TokenSecret = os.Getenv("TOKEN_SECRET")

	return config, nil
}
