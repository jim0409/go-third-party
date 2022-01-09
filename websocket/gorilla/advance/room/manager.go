package room

import (
	"hash/maphash"
	"log"
	"runtime"

	"github.com/gorilla/websocket"
)

type HubManager struct {
	hubs     []*Hub
	hashSeed maphash.Seed
	logger   loggerFunc
	findHdr  FindGroupMemberHdr
}

type FindGroupMemberHdr func(string, string) bool

func InitManager(findHdr FindGroupMemberHdr) *HubManager {
	// logger ... 內嵌 log.Printf ... 可以改成 utils 定義
	return &HubManager{
		hashSeed: maphash.MakeSeed(),
		logger:   log.Printf,
		findHdr:  findHdr,
	}
}

func (m *HubManager) NewHub() *Hub {
	return &Hub{
		register:       make(chan *WebSocketConn),
		unregister:     make(chan *WebSocketConn),
		broadcast:      make(chan *WebSocketEvent, broadcastQueueSize),
		stop:           make(chan struct{}),
		invalidateUser: make(chan string),
		directMessage:  make(chan *webSocketConnDirectMessage),
		logger:         m.logger,
	}
}

func (m *HubManager) HubStart() {
	// Total number of hubs is twice the number of CPUs.
	numberOfHubs := runtime.NumCPU() * 2
	m.logger("Starting webSocket hubs, amount: %d", numberOfHubs)

	hubs := make([]*Hub, numberOfHubs)

	for i := 0; i < numberOfHubs; i++ {
		hubs[i] = m.NewHub()
		hubs[i].connectionIndex = i
		hubs[i].Start()
	}

	// Assigning to the hubs slice without any mutex is fine because it is only assigned once
	// during the start of the program and always read from after that.
	m.hubs = hubs
}

func (m *HubManager) HubStop() {
	m.logger("Stopping webSocket hubs")

	for _, hub := range m.hubs {
		hub.Stop()
	}
}

// GetHubForUserID returns the hub for a given user id
func (m *HubManager) GetHubForUserID(userID string) *Hub {
	// TODO: check if caching the userID -> hub mapping
	// is worth the memory tradeoff.
	// https://mattermost.atlassian.net/browse/MM-26629.
	var hash maphash.Hash
	hash.SetSeed(m.hashSeed)
	hash.Write([]byte(userID))
	index := hash.Sum64() % uint64(len(m.hubs))

	return m.hubs[int(index)]
}

// HubRegister registers a connection to a hub.
func (m *HubManager) HubRegister(wc *WebSocketConn) {
	hub := m.GetHubForUserID(wc.UserID)
	if hub != nil {
		hub.Register(wc)
	}
}

// HubUnregister unregisters a connection from a hub.
func (m *HubManager) HubUnregister(wc *WebSocketConn) {
	hub := m.GetHubForUserID(wc.UserID)
	if hub != nil {
		hub.Unregister(wc)
	}
}

// Publish broadcast event to a hub with specific user or every hub.
func (m *HubManager) Publish(msg *WebSocketEvent) {
	// func (m *HubManager) Publish(msg WebSocketMessage) {
	// m.logger(message.Data)
	if err := msg.EnPack(); err != nil {
		m.logger("error happened while msg enpack!")
		return
	}

	if msg.Broadcast.UserID != "" {
		m.logger("enter here ... 1")
		hub := m.GetHubForUserID(msg.Broadcast.UserID)
		if hub != nil {
			m.logger("enter here ... 2")
			hub.Broadcast(msg)
		}
	} else {
		// m.logger("enter here ... 3")
		for _, hub := range m.hubs {
			// m.logger("enter here ... 4")
			hub.Broadcast(msg)
		}
	}
}

// InvalidateUser ban a user.
func (m *HubManager) InvalidateUser(userID string) {
	hub := m.GetHubForUserID(userID)
	if hub != nil {
		hub.InvalidateUser(userID)
	}
}

func (m *HubManager) NewWebSocketConn(ws *websocket.Conn, session Session) *WebSocketConn {
	// logger ... 內嵌 log.Printf ... 可以改成 utils 定義
	wc := &WebSocketConn{
		WebSocket:    ws,
		UserID:       session.UserID,
		send:         make(chan WebSocketMessage, sendQueueSize),
		pumpFinished: make(chan struct{}),
		manager:      m,
		logger:       log.Printf,
	}

	return wc
}
