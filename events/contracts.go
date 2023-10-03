package events

type EventMarshaler interface {
	Marshal(eventType EventType, eventData interface{}) ([]byte, error)
}
