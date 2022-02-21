package natsdk

import (
	"testing"
)

func TestPubToSubj(t *testing.T) {

	natsUrl := "nats://127.0.0.1:4222"
	sc, err := InitNatsServer("test-pub", natsUrl, "stan", "stan")
	if err != nil {
		panic(err)
	}
	defer sc.Closed()

	if err := sc.SendMsg("stan", []byte("test")); err != nil {
		panic(err)
	}
}
