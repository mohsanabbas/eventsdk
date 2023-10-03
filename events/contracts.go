package events

type EventMarshaler interface {
	Marshal(eventType EventType, eventData interface{}) ([]byte, error)
	MarshalBatch(events []Event) ([]byte, error)
}
