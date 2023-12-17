package util

import (
	"time"

	"github.com/spf13/viper"
)

// Config stores all the configuration of the application.
// The values are read by viper from a config file or environment variables.
type Config struct {
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	DBName               string        `mapstructure:"DB_NAME"`
	DBUsername           string        `mapstructure:"DB_USERNAME"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBHost               string        `mapstructure:"DB_HOST"`
	DBPort               string        `mapstructure:"DB_PORT"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
}

// LoadConfig reads configuration from config file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
