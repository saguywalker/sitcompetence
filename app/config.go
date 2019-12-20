package app

import (
	"crypto/sha256"
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
	secret := sha256.Sum256([]byte(viper.GetString("SecretKey")))
	config := &Config{
		SecretKey: secret[:],
	}

	if len(config.SecretKey) == 0 {
		return nil, errors.New("secretkey must not be empty")
	}

	return config, nil
}
