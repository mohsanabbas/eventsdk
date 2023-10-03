package eventsdk

import (
	"event-tracking-sdk/events"
	"event-tracking-sdk/transport"
)

type SDK struct {
	Endpoint   string
	Transport  transport.Transporter
	EventMarsh events.EventMarshaler
}

func New(endpoint string, trans transport.Transporter, marshaller events.EventMarshaler) *SDK {
	return &SDK{
		Endpoint:   endpoint,
		Transport:  trans,
		EventMarsh: marshaller,
	}
}

func (sdk *SDK) SendEvent(eventType events.EventType, eventData interface{}) error {

	data, err := sdk.EventMarsh.Marshal(eventType, eventData)
	if err != nil {
		return err
	}
	return sdk.Transport.Send(data, sdk.Endpoint)
}
