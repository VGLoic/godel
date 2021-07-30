package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.SetEnvPrefix("SERVER")
	viper.AutomaticEnv()

	fromDocker := viper.GetBool("FROM_DOCKER")

	if fromDocker {
		return nil
	}

	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}
