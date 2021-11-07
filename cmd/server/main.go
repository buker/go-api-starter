package main

import (
	// "fmt"

	"os"

	"github.com/buker/go-api-starter/internal/config"
	"github.com/buker/go-api-starter/internal/database"
	"github.com/buker/go-api-starter/internal/middleware"
	"github.com/buker/go-api-starter/internal/router"
	log "github.com/sirupsen/logrus"
)

var configure = config.Config()

func init() {
	if configure.Server.ServerEnv != "local" {
		log.SetFormatter(&log.JSONFormatter{})
	} else if configure.Server.ServerEnv == "local" {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	}
	log.SetReportCaller(true)

	log.SetOutput(os.Stdout)
	if configure.Logger.LogLevel == "debug" {
		log.SetLevel(log.DebugLevel)
	} else if configure.Logger.LogLevel == "info" {
		log.SetLevel(log.InfoLevel)
	} else if configure.Logger.LogLevel == "warn" {
		log.SetLevel(log.WarnLevel)
	} else if configure.Logger.LogLevel == "error" {
		log.SetLevel(log.ErrorLevel)
	} else if configure.Logger.LogLevel == "fatal" {
		log.SetLevel(log.FatalLevel)
	} else if configure.Logger.LogLevel == "panic" {
		log.SetLevel(log.PanicLevel)
	}
}

func main() {
	middleware.SentryInit(configure.Logger.SentryDsn)
	if err := database.InitDB().Error; err != nil {
		log.Error(err)
	}

	// JWT
	middleware.MySigningKey = []byte(configure.Server.ServerJWT.Key)
	middleware.JWTExpireTime = configure.Server.ServerJWT.Expire
	router := router.SetupRouter()
	log.Info("Server will start")
	err := router.Run(":" + configure.Server.ServerPort)
	if err != nil {
		log.Error(err)
	}
}
