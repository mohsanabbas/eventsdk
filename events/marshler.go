package events

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JSONEventMarshaler struct{}

// Marshal takes an individual event type and its data, validates the event, and then marshals it into a JSON byte slice.
// If the event is invalid or if there's an error during marshaling, an error is returned.
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

// MarshalBatch takes a slice of events, validates each event, and then marshals the entire batch into a JSON byte slice.
// If any event in the batch is invalid or if there's an error during marshaling, an error is returned.
func (j *JSONEventMarshaler) MarshalBatch(events []Event) ([]byte, error) {
	// A map to store error counts and sample offending event names
	errorSummary := make(map[string][]string)

	for i, event := range events {
		if !IsValidEventType(event.Type) {
			return nil, fmt.Errorf("invalid event type: %s", event.Type)
		}

		err := j.validateEvent(event.Type, event.Custom)
		if err != nil {
			errorKey := err.Error()
			if _, exists := errorSummary[errorKey]; !exists {
				errorSummary[errorKey] = make([]string, 0, 3)
			}

			// Only store up to 3 sample event names for each error type
			if len(errorSummary[errorKey]) < 3 {
				errorSummary[errorKey] = append(errorSummary[errorKey], fmt.Sprintf("%s-%d", event.Type, i))
			}
		}
	}

	if len(errorSummary) > 0 {
		messages := make([]string, 0, len(errorSummary))
		for err, samples := range errorSummary {
			if len(samples) == 3 {
				messages = append(messages, fmt.Sprintf("%s: %s ... (and more)", err, strings.Join(samples, ", ")))
			} else {
				messages = append(messages, fmt.Sprintf("%s: %s", err, strings.Join(samples, ", ")))
			}
		}
		return nil, fmt.Errorf("Validation errors:\n%s", strings.Join(messages, "\n"))
	}

	//data, err := json.MarshalIndent(events, "", "    ")
	//if err != nil {
	//	return nil, fmt.Errorf("failed to marshal events: %v", err)
	//}
	//fmt.Println(string(data))
	//return data, nil
	return json.Marshal(events)
}
