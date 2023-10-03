package eventsdk

import (
	"event-tracking-sdk/events"
	"time"
)

type Batch struct {
	sdk         *SDK
	EventBuffer []events.Event
	Timer       *time.Timer
	BufferSize  int
	MaxWaitTime time.Duration
}

func NewBatcher(sdk *SDK, bufferSize int, maxWaitTime time.Duration) *Batch {
	return &Batch{
		sdk:         sdk,
		BufferSize:  bufferSize,
		MaxWaitTime: maxWaitTime,
	}
}

func (b *Batch) AddEvent(eventType events.EventType, eventData interface{}) error {
	ev := events.Event{
		Type:   eventType,
		Custom: eventData,
	}

	b.EventBuffer = append(b.EventBuffer, ev)

	if len(b.EventBuffer) >= b.BufferSize {
		return b.flush()
	}

	if b.Timer == nil {
		b.Timer = time.AfterFunc(b.MaxWaitTime, func() {
			_ = b.flush()
		})
	}

	return nil
}

func (b *Batch) flush() error {
	if b.Timer != nil {
		b.Timer.Stop()
		b.Timer = nil
	}
	if len(b.EventBuffer) == 0 {
		return nil
	}

	err := b.sdk.SendBatch(b.EventBuffer)
	b.EventBuffer = nil
	return err
}
