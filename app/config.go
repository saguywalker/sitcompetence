package app

import (
	"errors"

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
		return nil, errors.New("secretkey must not be empty")
	}

	return config, nil
}
