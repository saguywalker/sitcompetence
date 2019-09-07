package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	// A secret string used for session cookies, passwords, etc.
	SecretKey []byte
}

// InitConfig returns config struct
func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(viper.GetString("SecretKey")),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("SecretKey must be set")
	}
	return config, nil
}
