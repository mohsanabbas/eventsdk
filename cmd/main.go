package main

import (
	eventsdk "event-tracking-sdk"
	"event-tracking-sdk/events"
	"event-tracking-sdk/transport"
	"fmt"
	"os"
	"time"
)

func main() {
	clientConfig := transport.HTTPClientConfig{
		SemaphoreBufferSize: 50,
		Timeout:             10 * time.Second,
		KeepAlive:           600 * time.Second,
		MaxIdleConns:        100,
		IdleConnTimeout:     90 * time.Second,
	}
	customClient := transport.NewCustomHTTPClient(clientConfig)

	httpTransport := &transport.HTTPTransport{
		Client: customClient,
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

	//// Without batching:
	//config := eventsdk.SDKConfig{
	//	EnableBatching: false,
	//}
	//sdk := eventsdk.New("https://httpbin.org/post", httpTransport, eventCodec, config)

	//trackEvent := events.Track{
	//	Event:        "Track",
	//	UserId:       "user123",
	//	AnonymousId:  "anon123",
	//	CreatedAt:    &time.Time{},
	//	Context:      map[string]interface{}{"landing": "raw"},
	//	Integrations: &events.Integrations{All: true, Mixpanel: true, Firebase: false, GoogleAnalytics: false, KinesisFirehose: false, Appsflyer: false, Slack: false},
	//	Properties: map[string]interface{}{
	//		"key_type":        "CONSUMER",
	//		"library":         "android",
	//		"library_version": "11.0.157",
	//		"mp_lib":          "event-tracking",
	//	},
	//}
	//*trackEvent.CreatedAt = time.Now()
	//err := sdk.SendEvent("track", trackEvent)
	//if err != nil {
	//	_, _ = fmt.Fprintf(os.Stderr, "%v", err)
	//}

	// For batching:
	config := eventsdk.SDKConfig{
		EnableBatching: true,
		BatchSize:      10,
		MaxWaitTime:    1 * time.Minute,
	}
	sdk := eventsdk.New("https://httpbin.org/post", httpTransport, eventCodec, config)
	// Batch process test
	for i := 0; i < 50; i++ {
		trackEvent := events.Track{
			Event:        fmt.Sprintf("Track-%d", i),
			UserId:       "user123",
			AnonymousId:  fmt.Sprintf("anon123-%d", i),
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

	// Sleep for a duration slightly greater than MaxWaitTime
	time.Sleep(1*time.Minute + 5*time.Second)

}
