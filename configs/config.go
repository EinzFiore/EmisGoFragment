package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AccountServiceUrl     string `mapstructure:"ACCOUNT_SERVICE_URL"`
	InstitutionServiceUrl string `mapstructure:"INSTITUTION_SERVICE_URL"`
	LogPath               string `mapstructure:"LOG_PATH"`
	MaxRequestTimeout     string `mapstructure:"MAX_REQUEST_TIMEOUT"`
}

// LoadConfig reads configuration from file or environment variables.
func loadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.SetDefault("MAX_REQUEST_TIMEOUT", "60")

	err = viper.Unmarshal(&config)
	return
}

func GetConfig() Config {
	config, err := loadConfig()
	if err != nil {
		// if file .env not found
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// set config from env os
			conf := Config{
				AccountServiceUrl:     os.Getenv("ACCOUNT_SERVICE_URL"),
				InstitutionServiceUrl: os.Getenv("INSTITUTION_SERVICE_URL"),
				LogPath:               os.Getenv("LOG_PATH"),
				MaxRequestTimeout:     os.Getenv("MAX_REQUEST_TIMEOUT"),
			}

			return conf
		} else {
			fmt.Println(err.Error())
		}
	}

	return config
}
