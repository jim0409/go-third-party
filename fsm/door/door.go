package door

import (
	"context"
	"fmt"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := &Door{
		To: to,
	}

	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"enter_state": func(_ctx context.Context, e *fsm.Event) { d.enterState(_ctx, e) },
		},
	)

	return d
}

func (d *Door) enterState(context context.Context, e *fsm.Event) {
	fmt.Printf("The door to %s is %s\n", d.To, e.Dst)
}
