package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string        `mapstructure:"DB_HOST"`
	DBPort         string        `mapstructure:"DB_PORT"`
	DBUsername     string        `mapstructure:"DB_USER"`
	DBPassword     string        `mapstructure:"DB_PASSWORD"`
	DBName         string        `mapstructure:"DB_NAME"`
	Port           int           `mapstructure:"PORT"`
	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRY"`
	TokenAge       int           `mapstructure:"TOKEN_MAXAGE"`
	TokenSecret    string        `mapstructure:"TOKEN_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
