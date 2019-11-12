package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config defines database configuration
type Config struct {
	DatabaseURI  string
	AwsAccessKey string
	AwsSecretKey string
	AwsBucket    string
	AwsRegion    string
	AwsPath      string
}

// InitConfig return reference to configuration
func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI:  viper.GetString("DatabaseURI"),
		AwsAccessKey: viper.GetString("AwsAccessKey"),
		AwsSecretKey: viper.GetString("AwsSecretKey"),
		AwsBucket:    viper.GetString("AwsBucket"),
		AwsRegion:    viper.GetString("AwsRegion"),
		AwsPath:      viper.GetString("AwsPath"),
	}
	if config.DatabaseURI == "" || config.AwsBucket == "" || config.AwsAccessKey == "" || config.AwsSecretKey == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}

	if config.AwsPath == "" {
		config.AwsPath = "/media/"
	}

	if config.AwsRegion == "" {
		config.AwsRegion = "ap-southeast-1"
	}

	return config, nil
}
