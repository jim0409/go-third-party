package main

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)

	fmt.Println(fsm.Current())

	err := fsm.Event(context.Background(), "open")
	// err := fsm.Event(fsm.Current(), "open")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())

	err = fsm.Event(context.Background(), "close")
	// err = fsm.Event(fsm.Current(), "close")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(fsm.Current())
}
