package events

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// validateEvent is a generic function that routes validation based on EventType.
func (j *JSONEventMarshaler) validateEvent(eventType EventType, eventData interface{}) error {

	switch eventType {
	case IdentifyEvent:
		identify, ok := eventData.(Identify)
		if !ok {
			return fmt.Errorf("invalid event data for %s event", IdentifyEvent)
		}
		return ValidateIdentify(identify)

	case TrackEvent:
		track, ok := eventData.(Track)
		if !ok {
			return fmt.Errorf("invalid event data for %s event", TrackEvent)
		}
		return ValidateTrack(track)

	// other event

	default:
		return fmt.Errorf("unrecognized event type: %s", eventType)
	}
}

func ValidateIdentify(identify Identify) error {
	if identify.Context == nil {
		return errors.New("context cannot be nil")
	}
	if identify.Integrations == nil {
		return errors.New("integrations cannot be nil")
	}

	return nil
}

func ValidateTrack(track Track) error {
	if track.Properties == nil || len(track.Properties) == 0 {
		return errors.New("properties cannot be empty")
	}
	if err := validateEventName(track.Event); err != nil {
		return err
	}
	return nil
}

func validateEventName(eventName string) error {
	if eventName == "" {
		return errors.New("event name cannot be empty")
	}

	words := strings.Split(eventName, " ")
	for _, word := range words {
		if !unicode.IsUpper(rune(word[0])) {
			return fmt.Errorf("each word in event name should start with a capital letter: %s", eventName)
		}
	}

	if strings.Contains(eventName, "  ") {
		return errors.New("event name should have single spaces between words")
	}

	return nil
}

var validEventTypes = map[EventType]bool{
	IdentifyEvent: true,
	TrackEvent:    true,
	PageEvent:     true,
	ScreenEvent:   true,
	GroupEvent:    true,
	AliasEvent:    true,
}

func IsValidEventType(eventType EventType) bool {
	_, valid := validEventTypes[eventType]
	return valid
}
