package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

// subscriber represents a message receiver
type subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

// chatServer enables broadcasting to a set of subscribers.
type chatServer struct {
	subscriberMessageBuffer int //  max number of messages that can be queued
	msghandler              func(f string, v ...interface{})
	serveMux                http.ServeMux
	subscribersMu           sync.Mutex
	subscribers             map[*subscriber]struct{}
}

func (cs *chatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cs.serveMux.ServeHTTP(w, r)
}

func newChatServer() *chatServer {
	cs := &chatServer{
		subscriberMessageBuffer: 16,
		msghandler:              log.Printf,
		subscribers:             make(map[*subscriber]struct{}),
	}
	cs.serveMux.Handle("/", http.FileServer(http.Dir("./js")))
	cs.serveMux.HandleFunc("/subscribe", cs.subscribeHandler)

	return cs
}

func (cs *chatServer) cronJob() {
	go func() {
		for {
			// TODO: handle instant connect
			time.Sleep(time.Second * 1)
			cs.publish([]byte("test"))
		}
	}()
}

// subscribeHandler accepts the WebSocket connection and messages.
func (cs *chatServer) subscribeHandler(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		cs.msghandler("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "")

	err = cs.subscribe(r.Context(), c)
	if errors.Is(err, context.Canceled) {
		return
	}
	if websocket.CloseStatus(err) == websocket.StatusNormalClosure ||
		websocket.CloseStatus(err) == websocket.StatusGoingAway {
		return
	}
	if err != nil {
		cs.msghandler("%v", err)
		return
	}
}

// subscribe creates a subscriber with a buffered msgs chan with CloseRead to keep reading from the connection
func (cs *chatServer) subscribe(ctx context.Context, c *websocket.Conn) error {
	ctx = c.CloseRead(ctx)

	s := &subscriber{
		msgs: make(chan []byte, cs.subscriberMessageBuffer),
		closeSlow: func() {
			c.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
		},
	}
	cs.addSubscriber(s)
	defer cs.deleteSubscriber(s)

	for {
		select {
		case msg := <-s.msgs:
			err := writeTimeout(ctx, time.Second*5, c, msg)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

// publish publishes the msg to all subscribers.
func (cs *chatServer) publish(msg []byte) {
	cs.subscribersMu.Lock()
	defer cs.subscribersMu.Unlock()

	for s := range cs.subscribers {
		select {
		case s.msgs <- msg:
		default:
			go s.closeSlow()
		}
	}
}

// addSubscriber registers a subscriber.
func (cs *chatServer) addSubscriber(s *subscriber) {
	cs.subscribersMu.Lock()
	cs.subscribers[s] = struct{}{}
	cs.subscribersMu.Unlock()
}

// deleteSubscriber deletes the given subscriber.
func (cs *chatServer) deleteSubscriber(s *subscriber) {
	cs.subscribersMu.Lock()
	delete(cs.subscribers, s)
	cs.subscribersMu.Unlock()
}

func writeTimeout(ctx context.Context, timeout time.Duration, c *websocket.Conn, msg []byte) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return c.Write(ctx, websocket.MessageText, msg)
}
