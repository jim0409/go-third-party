package room

import (
	"runtime/debug"
	"sync/atomic"
)

const (
	broadcastQueueSize = 4096
)

type webSocketConnDirectMessage struct {
	conn *WebSocketConn
	msg  *WebSocketDirectMessage
}

type Hub struct {
	connectionCount int64
	register        chan *WebSocketConn
	unregister      chan *WebSocketConn
	broadcast       chan *WebSocketEvent
	stop            chan struct{}
	invalidateUser  chan string
	directMessage   chan *webSocketConnDirectMessage
	explicitStop    bool
	connectionIndex int
	logger          loggerFunc
}

func (h *Hub) Start() {
	var doStart, doRecover, doRecoverableStart func()
	doRecoverableStart = func() {
		defer doRecover()
		doStart()
	}

	doRecover = func() {
		if h.explicitStop {
			return
		}

		if r := recover(); r != nil {
			h.logger("Recovering from Hub panic: %v", r)
		} else {
			h.logger("Hub stopped unexpectedly. Recovering.")
		}

		h.logger(string(debug.Stack()))
	}

	doStart = func() {

		// log.Println("start new hub connection Index")
		connIndex := newHubConnectionIndex()

		for {
			select {
			case webSocketConn := <-h.register:
				h.logger("append usr ... %v\n", webSocketConn.UserID)
				connIndex.Add(webSocketConn)
				atomic.StoreInt64(&h.connectionCount, int64(len(connIndex.All())))
				webSocketConn.send <- webSocketConn.createHelloMessage()
			case webSocketConn := <-h.unregister:
				connIndex.Remove(webSocketConn)
				atomic.StoreInt64(&h.connectionCount, int64(len(connIndex.All())))
				close(webSocketConn.send)
				if len(webSocketConn.UserID) == 0 {
					continue
				}
			case userID := <-h.invalidateUser:
				for _, webSocketConn := range connIndex.ForUser(userID) {
					go webSocketConn.Close()
				}
			case directMsg := <-h.directMessage:
				if !connIndex.Has(directMsg.conn) {
					continue
				}
				select {
				case directMsg.conn.send <- directMsg.msg:
				default:
					h.logger("broadcast: cannot send, closing webSocket for user: %s", directMsg.conn.UserID)
					close(directMsg.conn.send)
					connIndex.Remove(directMsg.conn)
				}
			case msg := <-h.broadcast:
				broadcast := func(webSocketConn *WebSocketConn) {
					h.logger("broadcast message, msg:%+v", msg)
					if !connIndex.Has(webSocketConn) {
						return
					}

					if webSocketConn.shouldSendEvent(msg) {
						select {
						case webSocketConn.send <- msg:
						default:
							h.logger("broadcast: cannot send, closing webSocket for user: %s", webSocketConn.UserID)
							close(webSocketConn.send)
							connIndex.Remove(webSocketConn)
						}
					}
				}

				// h.logger("test ...")
				if msg.Broadcast.UserID != "" {
					h.logger("enter here ... 5")
					target := connIndex.ForUser(msg.Broadcast.UserID)
					for _, webSocketConn := range target {
						broadcast(webSocketConn)
					}
				} else {
					// h.logger("enter here ... 6")
					target := connIndex.All()
					for webSocketConn := range target {
						h.logger("enter here ... 7")
						broadcast(webSocketConn)
					}
				}

			case <-h.stop:
				for webSocketConn := range connIndex.All() {
					webSocketConn.Close()
				}
				h.explicitStop = true
				return
			}
		}
	}

	go doRecoverableStart()
}

func (h *Hub) Stop() {
	close(h.stop)
}

// Register registers a connection to the hub.
func (h *Hub) Register(wc *WebSocketConn) {
	select {
	case h.register <- wc:
	case <-h.stop:
	}
}

// Unregister unregisters a connection from the hub.
func (h *Hub) Unregister(wc *WebSocketConn) {
	select {
	case h.unregister <- wc:
	case <-h.stop:
	}
}

// Broadcast broadcasts the message to all connections in the hub.
func (h *Hub) Broadcast(message *WebSocketEvent) {
	// func (h *Hub) Broadcast(message WebSocketMessage) {
	if message != nil {
		select {
		case h.broadcast <- message:
			// log.Println("broadcast!!")
		case <-h.stop:
		}
	}
}

// InvalidateUser invalidates the cache for the given user.
func (h *Hub) InvalidateUser(userID string) {
	select {
	case h.invalidateUser <- userID:
	case <-h.stop:
	}

}

// SendMessage sends the given message to the given connection.
// func (h *Hub) SendMessage(conn *WebSocketConn, msg WebSocketMessage) {
func (h *Hub) SendMessage(conn *WebSocketConn, msg *WebSocketDirectMessage) {
	select {
	case h.directMessage <- &webSocketConnDirectMessage{
		conn: conn,
		msg:  msg,
	}:
	case <-h.stop:
	}
}

func (h *Hub) GetConnectionCount() int64 {
	return atomic.LoadInt64(&h.connectionCount)
}

// hubConnectionIndex provides fast addition, removal, and iteration of web connections.
// It requires 3 functionalities which need to be very fast:
// - check if a connection exists or not.
// - get all connections for a given userID.
// - get all connections.
type hubConnectionIndex struct {
	// byUserID stores the list of connections for a given userID
	byUserID map[string][]*WebSocketConn
	// byConnection serves the dual purpose of storing the index of the webconn
	// in the value of byUserID map, and also to get all connections.
	byConnection map[*WebSocketConn]int
}

func newHubConnectionIndex() *hubConnectionIndex {
	return &hubConnectionIndex{
		byUserID:     make(map[string][]*WebSocketConn),
		byConnection: make(map[*WebSocketConn]int),
	}
}

func (i *hubConnectionIndex) Add(wc *WebSocketConn) {
	i.byUserID[wc.UserID] = append(i.byUserID[wc.UserID], wc)
	i.byConnection[wc] = len(i.byUserID[wc.UserID]) - 1
}

func (i *hubConnectionIndex) Remove(wc *WebSocketConn) {
	userConnIndex, ok := i.byConnection[wc]
	if !ok {
		return
	}

	// get the conn slice.
	userConnections := i.byUserID[wc.UserID]
	// get the last connection.
	last := userConnections[len(userConnections)-1]
	// set the slot that we are trying to remove to be the last connection.
	userConnections[userConnIndex] = last
	// remove the last connection from the slice.
	i.byUserID[wc.UserID] = userConnections[:len(userConnections)-1]
	// set the index of the connection that was moved to the new index.
	i.byConnection[last] = userConnIndex

	delete(i.byConnection, wc)
}

func (i *hubConnectionIndex) Has(wc *WebSocketConn) bool {
	_, ok := i.byConnection[wc]
	return ok
}

func (i *hubConnectionIndex) ForUser(id string) []*WebSocketConn {
	return i.byUserID[id]
}

func (i *hubConnectionIndex) All() map[*WebSocketConn]int {
	return i.byConnection
}
