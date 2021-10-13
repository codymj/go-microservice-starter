package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Registry for configuration values
var Registry *viper.Viper

// Set configurations
func Set() {
	Registry = viper.GetViper()
	Registry.AddConfigPath(".")
	Registry.AddConfigPath("../..")
	Registry.SetConfigFile("settings.yaml")

	err := Registry.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w \n", err))
	}
}
