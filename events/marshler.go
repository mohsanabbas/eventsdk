package events

import (
	"encoding/json"
	"fmt"
)

type JSONEventMarshaler struct{}

func (j *JSONEventMarshaler) Marshal(eventType EventType, eventData interface{}) ([]byte, error) {
	if !IsValidEventType(eventType) {
		return nil, fmt.Errorf("invalid event type: %s", eventType)
	}
	if err := j.validateEvent(eventType, eventData); err != nil {
		return nil, err
	}

	event := Event{
		Type:   eventType,
		Custom: eventData,
	}

	data, err := json.MarshalIndent(event, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal event: %v", err)
	}

	fmt.Println(string(data))
	return data, nil
}
