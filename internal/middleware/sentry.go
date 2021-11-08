package middleware

import (
	"github.com/onrik/logrus/sentry"
	log "github.com/sirupsen/logrus"
)
// SentryHook is a hook to send exceptions to an exception-tracking service compatible with the Sentry protocol.
func SentryInit(dsn string) {
	sentryHook, err := sentry.NewHook(sentry.Options{
		Dsn: dsn,
	}, log.PanicLevel, log.FatalLevel, log.ErrorLevel)
	if err != nil {
		log.Error(err)
		return
	}
	defer sentryHook.Flush()
	log.AddHook(sentryHook)
}
