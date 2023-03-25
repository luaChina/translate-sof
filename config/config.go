package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type EnvConfig struct {
	Host     string `mapstructure:"DB_HOST,omitempty"`
	Port     int    `mapstructure:"DB_PORT,omitempty"`
	User     string `mapstructure:"DB_USER,omitempty"`
	Password string `mapstructure:"DB_PASSWORD,omitempty"`
	ApiKey   string `mapstructure:"API_KEY,omitempty"`
	ApiProxy string `mapstructure:"API_PROXY,omitempty"`
}

var SecretConfig EnvConfig

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&SecretConfig); err != nil {
		panic(err)
	}
	fmt.Print(SecretConfig)
}
