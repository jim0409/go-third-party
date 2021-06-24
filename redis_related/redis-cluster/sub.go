package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/go-redis/redis"
)

type PubSub struct {
	client *redis.ClusterClient
}

var Service *PubSub

func init() {
	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{
			":7000",
			":7001",
			":7002",
			":7003",
			":7004",
			":7005",
		},
	})
	Service = &PubSub{client}
}

func console(channel, payload string) {
	log.Println(payload)
}

func main() {
	// Create subscriber
	_, err := NewSubscriber("secKey", console)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	log.Print("Subscriptions done. Publishing...")

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println()
		log.Println(sig)
		done <- true
	}()

	<-done
}

type Subscriber struct {
	pubsub   *redis.PubSub
	channel  string
	callback processFunc
}

type processFunc func(string, string)

func NewSubscriber(channel string, fn processFunc) (*Subscriber, error) {
	var err error
	// TODO Timeout param?

	s := Subscriber{
		pubsub:   Service.client.Subscribe(),
		channel:  channel,
		callback: fn,
	}

	// Subscribe to the channel
	err = s.subscribe()
	if err != nil {
		return nil, err
	}

	// Listen for messages
	go s.listen()

	return &s, nil
}

func (s *Subscriber) subscribe() error {
	var err error

	err = s.pubsub.Subscribe(s.channel)
	if err != nil {
		log.Println("Error subscribing to channel.")
		return err
	}
	return nil
}

func (s *Subscriber) listen() error {
	var channel string
	var payload string

	for {
		msg, err := s.pubsub.Receive() // no timeouts
		if err != nil {
			if reflect.TypeOf(err) == reflect.TypeOf(&net.OpError{}) &&
				reflect.TypeOf(err.(*net.OpError).Err).String() == "*net.timeoutError" {
				// Timeout, ignore
				continue
			}
			// Actual error
			log.Print("Error in ReceiveTimeout()", err)
		}

		channel = ""
		payload = ""

		switch m := msg.(type) {
		case *redis.Subscription:
			log.Printf("Subscription Message: %v to channel '%v'. %v total subscriptions.", m.Kind, m.Channel, m.Count)
			continue
		case *redis.Message:
			channel = m.Channel
			payload = m.Payload
		}

		// Process the message
		go s.callback(channel, payload)
	}
}
