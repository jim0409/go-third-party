package main

import (
	"log"
	"sync"

	"github.com/nats-io/nats.go"
)

func main() {
	// [begin subscribe_json]
	// nc, err := nats.Connect("demo.nats.io")
	// nc, err := nats.Connect(nats.DefaultURL)
	nc, err := nats.Connect("nats://192.168.51.110:4222")
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

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	if _, err := ec.Subscribe("updates", func(s *stock) {
		log.Printf("Stock: %s - Price: %v", s.Symbol, s.Price)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	// [end subscribe_json]
}
