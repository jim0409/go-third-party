package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	fmt.Println("wait for ctrl + c")
	natsUrl := "nats://127.0.0.1:4222"
	sc, err := InitNatsServer("test-sub", natsUrl, "stan", "stan")
	if err != nil {
		sc.Closed()
	}

	// if err := sc.SendMsg("stan", []byte("test")); err != nil {
	// 	panic(err)
	// }
	sub, err := sc.RecvMsg("stan")
	if err != nil {
		sub.Unsubscribe()
		panic(err)
	}

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
