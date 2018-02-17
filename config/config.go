package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

// Init is an exported method that takes the environment starts the viper
// (external lib) and returns the configuration struct.
func Init(env string) {
	var err error
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName(env)

	if env == "db_seed" {
		v.AddConfigPath("../config/")
	}

	if env == "test" {
		v.AddConfigPath("../config/")
	} else {
		v.AddConfigPath("config/")
	}
	err = v.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}
	config = v
}

//GetConfig returns the config object
func GetConfig() *viper.Viper {
	return config
}
