package api

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config struct
type Config struct {
	IPAddress string

	// The port to bind the web application server to
	Port int

	// The number of proxies positioned in front of the API. This is used to interpret
	// X-Forwarded-For headers.
	// ProxyCount int

	Peers []string
}

// InitConfig return config struct
func InitConfig() (*Config, error) {
	config := &Config{
		Port: viper.GetInt("Port"),
		// ProxyCount: viper.GetInt("ProxyCount"),
		Peers: viper.GetStringSlice("Peers"),
		IPAddress: viper.GetString("BaseIP"),
	}
	if config.Port == 0 {
		config.Port = 3000
	}

	if len(config.Peers) == 0 {
		return nil, fmt.Errorf("Blockchain should have at least 1 peer")
	}
	return config, nil
}
