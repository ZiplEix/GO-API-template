package config

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	PORT        string `mapstructure:"PORT"`
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
			PORT:        os.Getenv("PORT"),
			JWT_SECRET:  os.Getenv("JWT_SECRET"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config
	if config.DB_USER == "" {
		err = errors.New("DB_USER is not set")
		return
	}
	if config.DB_PASSWORD == "" {
		err = errors.New("DB_PASSWORD is not set")
		return
	}
	if config.DB_NAME == "" {
		err = errors.New("DB_NAME is not set")
		return
	}
	if config.PORT == "" {
		err = errors.New("PORT is not set")
		return
	}
	if config.JWT_SECRET == "" {
		err = errors.New("JWT_SECRET is not set")
	}

	return
}
