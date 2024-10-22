package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost                string        `mapstructure:"DB_HOST"`
	DBPort                string        `mapstructure:"DB_PORT"`
	DBUsername            string        `mapstructure:"DB_USER"`
	DBPassword            string        `mapstructure:"DB_PASSWORD"`
	DBName                string        `mapstructure:"DB_NAME"`
	Port                  int           `mapstructure:"PORT"`
	RefreshTokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRY"`
	TokenExpiresIn        time.Duration `mapstructure:"TOKEN_EXPIRY"`
	TokenAge              int           `mapstructure:"TOKEN_MAXAGE"`
	TokenSecret           string        `mapstructure:"TOKEN_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.SetDefault("PORT", 4055)
	viper.SetDefault("TOKEN_MAXAGE", 3600)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Warning: No config file found. Relying on environment variables.\n")
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into struct, %w", err)
	}

	return config, nil
}
