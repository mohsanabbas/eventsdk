package main

import (
	eventsdk "event-tracking-sdk"
	"event-tracking-sdk/events"
	"event-tracking-sdk/transport"
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	httpTransport := &transport.HTTPTransport{
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
		RetryConf: transport.RetryConfig{
			MaxRetries:      3,
			InitialDelay:    500 * time.Millisecond,
			MaxDelay:        8 * time.Second,
			ExponentialBase: 2.0,
			Jitter:          true,
			RetryOnStatus:   []int{500, 502, 503, 504},
		},
	}

	eventCodec := &events.JSONEventMarshaler{}

	sdk := eventsdk.New("https://httpbin.org/post", httpTransport, eventCodec)

	trackEvent := events.Track{
		Event:        "Track",
		UserId:       "user123",
		AnonymousId:  "anon123",
		CreatedAt:    &time.Time{},
		Context:      map[string]interface{}{"landing": "raw"},
		Integrations: &events.Integrations{All: true, Mixpanel: true, Firebase: false, GoogleAnalytics: false, KinesisFirehose: false, Appsflyer: false, Slack: false},
		Properties: map[string]interface{}{
			"key_type":        "CONSUMER",
			"library":         "android",
			"library_version": "11.0.157",
			"mp_lib":          "event-tracking",
		},
	}
	*trackEvent.CreatedAt = time.Now()

	err := sdk.SendEvent("track", trackEvent)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v", err)
	}
}
