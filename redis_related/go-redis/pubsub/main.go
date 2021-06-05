package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/redis.v2"
)

type food struct {
	Name     string
	Calories float64
}

type car struct {
	Make  string
	Model string
}

func eat(channel, payload string) {
	var f food

	err := json.Unmarshal([]byte(payload), &f)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	log.Printf("Eating a %v.", f.Name)
}

func subChFn(channel, payload string) {
	// log.Println(payload)
	var c car

	err := json.Unmarshal([]byte(payload), &c)
	if err != nil {
		log.Printf("Unmarshal error: %v", err)
	}

	// publish cars "{\"Make\":\"Tesla\",\"Model\":\"Model S\"}"
	log.Printf("Driving a %v.", c.Make)
}

func main() {
	var pub *redis.IntCmd
	var err error

	// Create a subscriber
	// _, err = NewSubscriber("food", eat)
	// if err != nil {
	// 	log.Println("NewSubscriber() error", err)
	// }
	// -- Publish some stuf --
	// pub = Service.Publish("food", food{"Pizza", 50.1})
	// if err = pub.Err(); err != nil {
	// 	log.Print("PublishString() error", err)
	// }

	// pub = Service.Publish("food", food{"Big Mac", 200})
	// if err = pub.Err(); err != nil {
	// 	log.Print("PublishString() error", err)
	// }

	// Create another subscriber
	_, err = NewSubscriber("cars", subChFn)
	if err != nil {
		log.Println("NewSubscriber() error", err)
	}
	log.Print("Subscriptions done. Publishing...")

	pub = Service.Publish("cars", car{"Subaru", "Impreza"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}

	pub = Service.Publish("cars", car{"Tesla", "Model S"})
	if err = pub.Err(); err != nil {
		log.Print("PublishString() error", err)
	}
	_ = pub

	log.Print("Publishing done. Sleeping...")

	GracefulShutdown()
}

// gracefulShutdown: handle the worker connection
func GracefulShutdown() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// Log.Info("awaiting signal")
	log.Println("awaiting signal")
	<-done
	// Log.Info("exiting")
	log.Println("exiting")
}
