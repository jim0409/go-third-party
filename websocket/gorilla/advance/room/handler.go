package room

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
)

type WebSocketManagerIface interface {
	GetHubForUserID(string) *Hub
	Publish(*WebSocketEvent)
	// Publish(WebSocketMessage)
	HubRegister(*WebSocketConn)

	NewWebSocketConn(*websocket.Conn, Session) *WebSocketConn
	InvalidateUser(userID string)
}

const (
	sendQueueSize = 256
	bufferCap     = 2 * 1024
	writeWaitTime = 30 * time.Second
	pingWaitTime  = 2 * time.Minute
)

type loggerFunc func(string, ...interface{})

type WebSocketConn struct {
	WebSocket *websocket.Conn
	UserID    string

	send         chan WebSocketMessage
	pumpFinished chan struct{}
	logger       loggerFunc
	manager      *HubManager
}

// Close closes the WebConn.
func (wc *WebSocketConn) Close() {
	wc.WebSocket.Close()
	<-wc.pumpFinished
}

func (wc *WebSocketConn) shouldSendEvent(msg *WebSocketEvent) bool {
	isOmit := false
	isInGroup := true

	if msg.Broadcast.UserID != "" {
		return wc.UserID == msg.Broadcast.UserID
	}

	if len(msg.Broadcast.OmitUsers) > 0 {
		if _, ok := msg.Broadcast.OmitUsers[wc.UserID]; ok {
			isOmit = true
		}
	}

	if msg.Broadcast.GroupID != "" {
		wc.logger("__ msg.Broadcast.GroupID: %v __ usrID %v\n", msg.Broadcast.GroupID, wc.UserID)
		isInGroup = wc.manager.findHdr(msg.Broadcast.GroupID, wc.UserID)
	}

	return !isOmit && isInGroup
}

func (wc *WebSocketConn) createHelloMessage() *WebSocketEvent {
	msg := NewWebSocketEvent(WebsocketEventHello, "", wc.UserID, nil)
	return msg
}

func (wc *WebSocketConn) Pump() {
	go wc.writePump()
	go wc.readPump()
}

func (wc *WebSocketConn) writePump() {
	// ticker := time.NewTicker(pingInterval)

	defer func() {
		// ticker.Stop()
		wc.WebSocket.Close()
	}()

	var buf bytes.Buffer
	buf.Grow(bufferCap)
	enc := json.NewEncoder(&buf)

	for {
		select {
		case msg, ok := <-wc.send:
			wc.WebSocket.SetWriteDeadline(time.Now().Add(writeWaitTime))
			if !ok {
				// The hub closed the channel.
				wc.WebSocket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			buf.Reset()
			var err error
			evt, evtOk := msg.(*WebSocketEvent)
			if evtOk {
				err = evt.Encode(enc)
			} else {
				err = enc.Encode(msg)
			}
			if err != nil {
				wc.logger("Error in encoding websocket message, err: %s", err.Error())
				continue
			}

			if err := wc.WebSocket.WriteMessage(websocket.TextMessage, buf.Bytes()); err != nil {
				wc.logSocketErr("websocket.send", err)
				return
			}
		}
	}
}

func (wc *WebSocketConn) readPump() {
	defer func() {
		wc.manager.HubUnregister(wc)
		wc.WebSocket.Close()
		close(wc.pumpFinished)
	}()

	wc.WebSocket.SetReadLimit(MaxMessageSize)
	wc.WebSocket.SetReadDeadline(time.Now().Add(pingWaitTime))

	for {
		var req WebSocketEvent
		if err := wc.WebSocket.ReadJSON(&req); err != nil {
			wc.logSocketErr("websocket.read", err)
			return
		}
		if req.Event == WebsocketEventPing {
			wc.WebSocket.SetReadDeadline(time.Now().Add(pingWaitTime))
			wc.sendPong()
		}
	}
}

func (wc *WebSocketConn) sendPong() {
	hub := wc.manager.GetHubForUserID(wc.UserID)
	if hub == nil {
		return
	}

	data := map[string]interface{}{}
	data["time"] = GetMillis()
	msg := NewWebSocketDirectMessage(WebsocketEventPong, data)
	hub.SendMessage(wc, msg)
}

func (wc *WebSocketConn) logSocketErr(source string, err error) {
	if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseNoStatusReceived) {
		wc.logger("%s : client side closed socket", source)
	} else {
		wc.logger("%s : closing websocket, err: %s", source, err.Error())
	}
}
