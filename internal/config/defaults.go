package config

import "github.com/spf13/viper"

func Defaults() {
	viper.SetDefault("service.name", "")

	viper.SetDefault("vault.address", "")
	viper.SetDefault("vault.token", "")
	viper.SetDefault("vault.path", "")
}
