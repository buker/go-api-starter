package main

import (
	// "fmt"

	"os"

	"github.com/buker/go-api-starter/internal/config"
	"github.com/buker/go-api-starter/internal/middlewares"
	"github.com/buker/go-api-starter/internal/router"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configure = config.GetConfig()

func init() {
	if err := config.Setup(); err != nil {
		log.WithError(err).Fatal("Failed to setup configuration")
	}
	if viper.GetString("app.env") != "local" {
		log.SetFormatter(&log.JSONFormatter{})
	} else if viper.GetString("app.env") == "local" {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}
	log.SetReportCaller(true)

	log.SetOutput(os.Stdout)
	if viper.GetString("logger.logLevel") == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if viper.GetString("logger.logLevel") == "info" {
		log.SetLevel(log.InfoLevel)
	} else if viper.GetString("logger.logLevel") == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if viper.GetString("logger.logLevel") == "error" {
		log.SetLevel(log.ErrorLevel)
	} else if viper.GetString("logger.logLevel") == "fatal" {
		log.SetLevel(log.FatalLevel)
	} else if viper.GetString("logger.logLevel") == "panic" {
		log.SetLevel(log.PanicLevel)
	}
}

func main() {
	log.Info(viper.GetString("app.env"))
	log.Info(viper.GetString("app.env"))
	if err := config.Setup(); err != nil {
		log.WithError(err).Fatal("Failed to setup configuration")
	}
	middlewares.SentryInit(viper.GetString("logger.sentryDsn"))

	router := router.SetupRouter()
	log.Info("Server will start")
	err := router.Run(":" + viper.GetString("app.port"))
	if err != nil {
		log.Error(err)
	}
}
