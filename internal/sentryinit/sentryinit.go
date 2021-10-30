package sentryinit

import (
	"net/http"

	sentry "github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

func Init() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://a67153b6c1214429846bd148ec2e5be5@o380765.ingest.sentry.io/6004421",
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			if hint.Context != nil {
				if req, ok := hint.Context.Value(sentry.RequestContextKey).(*http.Request); ok {
					// You have access to the original Request here
					log.Info("Sentry request: %s", req)
				}
			}
			return event
		},
	}); err != nil {
		log.Fatalf("Sentry initialization failed: %v\n", err)
	}
}
