package room

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	WebsocketEventHello                  = "hello"
	WebsocketEventTyping                 = "typing"
	WebsocketEventMessage                = "message"
	WebsocketEventMemberJoin             = "member_join"
	WebsocketEventMemberLeft             = "member_left"
	WebsocketEventGroupAdd               = "group_add"
	WebsocketEventGroupLeft              = "group_left"
	WebsocketEventGroupRead              = "group_read"
	WebsocketEventGroupChangeDisplayName = "group_displayname"
	WebsocketEventGroupChangeIcon        = "group_icon"

	WebsocketEventPong = "pong"
	WebsocketEventPing = "ping"
)

const (
	MaxMessageSize = 8 * 1024
)

type Session struct {
	Token        string `json:"token" db:"token"`
	RefreshToken string `json:"refresh_token" db:"refresh_token"`
	Type         int    `json:"type" db:"type"`
	UserID       string `json:"user_id" db:"user_id"`
	DeviceID     string `json:"device_id" db:"device_id"`
	System       int    `json:"system" db:"system"`
	DeviceToken  string `json:"device_token" db:"device_token"`
	CreateAt     int64  `json:"create_at" db:"create_at"`
	ExpireAt     int64  `json:"expire_at" db:"expire_at"`
}

type WebSocketMessage interface {
	ToJson() string
}

type precomputedWebSocketEventJSON struct {
	Event     json.RawMessage
	Data      json.RawMessage
	Broadcast json.RawMessage
}

type WebSocketEvent struct {
	Event           string              `json:"event"`
	Data            interface{}         `json:"data,omitempty"`
	Broadcast       *WebSocketBroadcast `json:"-"`
	precomputedJSON *precomputedWebSocketEventJSON
}

type WebSocketBroadcast struct {
	UserID    string              `json:"user_id"`
	OmitUsers map[string]struct{} `json:"omit_users"`
	GroupID   string              `json:"group_id"`
}

func (ev *WebSocketEvent) EnPack() error {
	event, _ := json.Marshal(ev.Event)
	data, _ := json.Marshal(ev.Data)
	broadcast, _ := json.Marshal(ev.Broadcast)

	ev.precomputedJSON = &precomputedWebSocketEventJSON{
		Event:     json.RawMessage(event),
		Data:      json.RawMessage(data),
		Broadcast: json.RawMessage(broadcast),
	}

	return nil
}

func (ev *WebSocketEvent) ToJson() string {
	if ev.precomputedJSON != nil {
		return fmt.Sprintf(`{"event": %s, "data": %s}`, ev.precomputedJSON.Event, ev.precomputedJSON.Data)
	}
	j, err := json.Marshal(ev)

	if err != nil {
		return string(j)
	}
	return ""
}

func NewWebSocketEvent(event, groupID, userID string, omitUsers map[string]struct{}) *WebSocketEvent {
	// func NewWebSocketEvent(event, groupID, userID string, omitUsers map[string]struct{}) WebSocketMessage {
	return &WebSocketEvent{
		Event:     event,
		Broadcast: &WebSocketBroadcast{GroupID: groupID, UserID: userID, OmitUsers: omitUsers},
	}

}

// Encode encodes the event to the given encoder.
func (ev *WebSocketEvent) Encode(enc *json.Encoder) error {
	if ev.precomputedJSON != nil {
		return enc.Encode(json.RawMessage(
			fmt.Sprintf(`{"event": %s, "data": %s}`, ev.precomputedJSON.Event, ev.precomputedJSON.Data),
		))

	}

	return enc.Encode(ev)
}

func (ev *WebSocketEvent) GetEventType() string {
	return ev.Event
}

type WebSocketDirectMessage struct {
	Event string                 `json:"event"`
	Data  map[string]interface{} `json:"data"`
}

func NewWebSocketDirectMessage(event string, data map[string]interface{}) *WebSocketDirectMessage {
	// func NewWebSocketDirectMessage(event string, data map[string]interface{}) *WebSocketEvent {
	return &WebSocketDirectMessage{
		// return &WebSocketEvent{
		Event: event,
		Data:  data,
	}
}

// func (m *WebSocketDirectMessage) Add(key string, value interface{}) {
// 	m.Data[key] = value
// }

// func (m *WebSocketDirectMessage) GetEventType() string {
// 	return m.Event
// }

func (m *WebSocketDirectMessage) ToJson() string {
	if b, err := json.Marshal(m); err != nil {
		return ""
	} else {
		return string(b)
	}
}

// GetMillis is a convenience method to get milliseconds since epoch.
func GetMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
