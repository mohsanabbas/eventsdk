package eventsdk

import (
	"event-tracking-sdk/events"
	"event-tracking-sdk/transport"
	"time"
)

type SDKConfig struct {
	EnableBatching bool
	BatchSize      int
	MaxWaitTime    time.Duration
}

type SDK struct {
	Endpoint   string
	Transport  transport.Transporter
	EventMarsh events.EventMarshaler
	batch      *Batch
}

func New(endpoint string, trans transport.Transporter, marshaller events.EventMarshaler, config SDKConfig) *SDK {
	sdk := &SDK{
		Endpoint:   endpoint,
		Transport:  trans,
		EventMarsh: marshaller,
	}
	if config.EnableBatching {
		sdk.batch = NewBatcher(sdk, config.BatchSize, config.MaxWaitTime)
	}
	return sdk
}

func (sdk *SDK) SendEvent(eventType events.EventType, eventData interface{}) error {
	if sdk.batch != nil {
		return sdk.batch.AddEvent(eventType, eventData)
	}

	data, err := sdk.EventMarsh.Marshal(eventType, eventData)
	if err != nil {
		return err
	}
	return sdk.Transport.Send(data, sdk.Endpoint)
}

func (sdk *SDK) SendBatch(events []events.Event) error {
	data, err := sdk.EventMarsh.MarshalBatch(events)
	if err != nil {
		return err
	}
	return sdk.Transport.Send(data, sdk.Endpoint)
}
