package app

import "github.com/spf13/viper"

type Config struct {
	Port    int    `mapstructure:"port"`
	AppName string `mapstructure:"app-name"`
}

func LoadConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("env")
	viper.SetConfigType("json")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	return &config
}
