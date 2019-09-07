package db

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config defines database configuration
type Config struct {
	DatabaseURI string
}

// InitConfig return reference to configuration
func InitConfig() (*Config, error) {
	config := &Config{
		DatabaseURI: viper.GetString("DatabaseURI"),
	}
	if config.DatabaseURI == "" {
		return nil, fmt.Errorf("DatabaseURI must be set")
	}
	return config, nil
}
