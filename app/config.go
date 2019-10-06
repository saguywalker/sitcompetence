package app

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	// A secret string used for session cookies, passwords, etc.
	SecretKey []byte
	Username  string
	Password  string
}

// InitConfig returns config struct
func InitConfig() (*Config, error) {
	config := &Config{
		SecretKey: []byte(viper.GetString("SecretKey")),
		// Username:  viper.GetString("LdapUser"),
		// Password:  viper.GetString("LdapPassword"),
	}
	if len(config.SecretKey) == 0 {
		return nil, fmt.Errorf("SecretKey, LdapUsername and Password must be set")
	}

	return config, nil
}
