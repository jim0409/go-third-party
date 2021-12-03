package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

// subscriber represents a subscriber.
// Messages are sent on the msgs channel and if the client
// cannot keep up with the messages, closeSlow is called.
type subscriber struct {
	msgs      chan []byte
	closeSlow func()
}

// chatServer enables broadcasting to a set of subscribers.
type chatServer struct {
	// controls the max number of messages that can be queued for a subscriber
	subscriberMessageBuffer int

	// logf controls where logs are sent .. Defaults to log.Printf.
	logf func(f string, v ...interface{})

	// serveMux routes the various endpoints to the appropriate handler.
	serveMux http.ServeMux

	subscribersMu sync.Mutex
	subscribers   map[*subscriber]struct{}
}

func (cs *chatServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cs.serveMux.ServeHTTP(w, r)
}

func newChatServer() *chatServer {
	cs := &chatServer{
		subscriberMessageBuffer: 16,
		logf:                    log.Printf,
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
