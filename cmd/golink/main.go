// The golink server runs http://go/, a private shortlink service for tailnets.
package main

import (
	_ "embed"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/tailscale/golink"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(2* time.Second)
	defer sentry.Recover()
	if err := golink.Run(); err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
}
