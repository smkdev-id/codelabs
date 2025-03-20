package configs

import (
	"github.com/spf13/viper"
)

func LoadViperEnv() {
	viper.AddConfigPath("/")
	viper.SetConfigFile("env.yaml")
	viper.AutomaticEnv()

	if EnvException := viper.ReadInConfig(); EnvException != nil {
		panic(EnvException)
	}
}
