package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin publish_json]
	// nc, err := nats.Connect("demo.nats.io")
	// nc, err := nats.Connect(nats.DefaultURL)
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type stock struct {
		Symbol string
		Price  int
	}

	fmt.Println(time.Now())
	// Publish the message
	for i := 0; i < 100000; i++ {
		// if err := ec.Publish("updates", &stock{Symbol: "GOOG", Price: 1200}); err != nil {
		if err := ec.Publish("updates", &stock{Symbol: "GOOG", Price: i}); err != nil {
			log.Fatal(err)
		}
	}
	// [end publish_json]
	fmt.Println(time.Now())
}
