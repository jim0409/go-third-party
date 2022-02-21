package natsdk

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"testing"
)

func TestSub(t *testing.T) {

	fmt.Println("wait for ctrl + c")
	natsUrl := "nats://127.0.0.1:4222"
	sc, err := InitNatsServer("test-sub", natsUrl, "stan", "stan")
	if err != nil {
		panic(err)
	}
	defer sc.Closed()

	sub, err := sc.RecvMsg("stan")
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
