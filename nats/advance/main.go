package main

import (
	"go-third-party/nats/advance/natsdk"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// TODO:
/*
1. receive packets from stan
2. parse packets
3. write content to arangodb
*/

func main() {

	natsUrl := "nats://3.112.66.62:4222"
	sc, err := natsdk.InitNatsServer("stan-local", natsUrl, "stan", "stan")
	if err != nil {
		panic(err)
	}
	defer sc.Closed()

	sub, err := sc.RecvMsg("MatchUpdate")
	if err != nil {
		panic(err)
	}
	defer sub.Unsubscribe()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("%v\n", sig)

		done <- true
	}()

	<-done

}
