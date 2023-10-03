package events

import (
	"encoding/json"
	"time"
)

type KeyType string

const (
	KeyTypeConsumer KeyType = "CONSUMER"
	KeyTypeEmail    KeyType = "EMAIL"
	KeyTypeSeller   KeyType = "SELLER"
)

type EventType string

const (
	IdentifyEvent EventType = "identify"
	TrackEvent    EventType = "track"
	PageEvent     EventType = "page"
	ScreenEvent   EventType = "screen"
	GroupEvent    EventType = "group"
	AliasEvent    EventType = "alias"
)

type Alias struct {
	UserId       string                 `json:"userId"`
	PreviousId   string                 `json:"previousId"`
	KeyType      KeyType                `json:"keyType,omitempty"`
	Integrations *Integrations          `json:"integrations"`
	Context      map[string]interface{} `json:"context"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Alias) MarshalJSON() ([]byte, error) {
	type Alias_ Alias

	return json.Marshal(&struct {
		Alias_
		CreatedAt string `json:"createdAt"`
	}{
		Alias_:    (Alias_)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Identify struct {
	UserId       string                 `json:"userId"`
	AnonymousId  string                 `json:"anonymousId"`
	KeyType      KeyType                `json:"keyType"`
	Integrations *Integrations          `json:"integrations"`
	Traits       map[string]interface{} `json:"traits"`
	Context      map[string]interface{} `json:"context"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Identify) MarshalJSON() ([]byte, error) {
	type Alias Identify

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt,omitempty"`
	}{
		Alias:     (Alias)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Page struct {
	UserId       string                 `json:"userId"`
	AnonymousId  string                 `json:"anonymousId"`
	KeyType      KeyType                `json:"keyType,omitempty"`
	Integrations *Integrations          `json:"integrations"`
	Name         string                 `json:"name"`
	Context      map[string]interface{} `json:"context"`
	Properties   map[string]interface{} `json:"properties"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Page) MarshalJSON() ([]byte, error) {
	type Alias Page

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt"`
	}{
		Alias:     (Alias)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Screen struct {
	UserId       string                 `json:"userId"`
	AnonymousId  string                 `json:"anonymousId"`
	KeyType      KeyType                `json:"keyType"`
	Integrations *Integrations          `json:"integrations"`
	Name         string                 `json:"name"`
	Context      map[string]interface{} `json:"context"`
	Properties   map[string]interface{} `json:"properties"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Screen) MarshalJSON() ([]byte, error) {
	type Alias Screen

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt"`
	}{
		Alias:     (Alias)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Track struct {
	UserId       string                 `json:"userId"`
	AnonymousId  string                 `json:"anonymousId"`
	KeyType      KeyType                `json:"keyType"`
	Integrations *Integrations          `json:"integrations"`
	Event        string                 `json:"event"`
	Properties   map[string]interface{} `json:"properties"`
	Context      map[string]interface{} `json:"context"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Track) MarshalJSON() ([]byte, error) {
	type Alias Track

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt"`
	}{
		Alias:     (Alias)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Group struct {
	UserId       string                 `json:"userId"`
	AnonymousId  string                 `json:"anonymousId"`
	KeyType      KeyType                `json:"keyType,omitempty"`
	Integrations *Integrations          `json:"integrations"`
	GroupId      string                 `json:"groupId"`
	Context      map[string]interface{} `json:"context"`
	Traits       map[string]interface{} `json:"traits"`
	Properties   map[string]interface{} `json:"properties"`
	CreatedAt    *time.Time             `json:"createdAt"`
}

func (s Group) MarshalJSON() ([]byte, error) {
	type Alias Group

	return json.Marshal(&struct {
		Alias
		CreatedAt string `json:"createdAt"`
	}{
		Alias:     (Alias)(s),
		CreatedAt: s.CreatedAt.Format(DateTimeFormat),
	})
}

type Event struct {
	Type   EventType   `json:"type"`
	Custom interface{} `json:"custom"`
}

type BatchEvents struct {
	Batch []Event `json:"batch,omitempty"`
}

type Integrations struct {
	All             bool `json:"all,omitempty"`
	Mixpanel        bool `json:"mixpanel,omitempty"`
	Firebase        bool `json:"firebase,omitempty"`
	GoogleAnalytics bool `json:"googleAnalytics,omitempty"`
	KinesisFirehose bool `json:"kinesisFirehose,omitempty"`
	Appsflyer       bool `json:"appsflyer,omitempty"`
	Slack           bool `json:"slack,omitempty"`
}

const (
	DateTimeFormat = "2006-01-02T15:04:05.000-0700"
)
