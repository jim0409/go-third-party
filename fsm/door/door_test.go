package door

import (
	"context"
	"os"
	"testing"

	"github.com/looplab/fsm"
	"github.com/stretchr/testify/assert"
)

func TestDoorFsm(t *testing.T) {
	door := NewDoor("heaven")

	err := door.FSM.Event(context.Background(), "open")
	assert.Nil(t, err)
	assert.Equal(t, "open", door.FSM.Current())

	err = door.FSM.Event(context.Background(), "close")
	assert.Nil(t, err)
	assert.Equal(t, "closed", door.FSM.Current())
}

func TestFsmVisual(t *testing.T) {
	door := NewDoor("heaven")

	pic, err := fsm.VisualizeForMermaidWithGraphType(door.FSM, fsm.FlowChart)
	assert.Nil(t, err)

	res := `graph LR
    id0[closed]
    id1[open]

    id0 --> |open| id1
    id1 --> |close| id0

    style id0 fill:#00AA00
`

	assert.Equal(t, res, pic)
}

func BenchmarkDoorFsm(b *testing.B) {
	// ignore output
	os.Stdout = nil

	door := NewDoor("heaven")
	for i := 0; i < b.N; i++ {
		if door.FSM.Can("open") {
			door.FSM.Event(context.Background(), "open")
		}
		if door.FSM.Can("close") {
			door.FSM.Event(context.Background(), "close")
		}
	}
}
