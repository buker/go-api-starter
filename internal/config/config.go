package config

import (
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Struct for configuration
type Configuration struct {
}

var Config *Configuration

//Function create configuration
func Setup() error {
	var configuration *Configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.WithError(err).Panic("Error reading config file")
		return err
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.WithError(err).Panic("Unable to decode into struct")
		return err
	}
	Config = configuration
	return nil
}
func GetConfig() *Configuration {
	return Config
}
