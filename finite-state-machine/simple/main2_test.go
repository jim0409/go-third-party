package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. 先宣告一個 struct 並且實踐對應的方法
type testStateMachine struct {
	Login      bool
	Processing int // init/ run/ pause = 0, 1, -1
}

// 2. 根據 struct 實踐 StateMachineCallback 並且判斷回傳方法
func (sm *testStateMachine) StateMachineCallback(action string, args []interface{}) {
	switch action {
	case "login":
		sm.Login = true
	case "logout":
		sm.Login = false
	case "run":
		sm.Processing = 1
	case "pause":
		sm.Processing = -1
	}
}

// 3. 測試~~
func TestSecondStateMachine(t *testing.T) {
	var fsm testStateMachine
	tm := NewStateMachine(&fsm,
		Transition{From: "logout", Event: "login", To: "run", Action: "login"},
		Transition{From: "login", Event: "run", To: "pause", Action: "run"},
		Transition{From: "run", Event: "pause", To: "logout", Action: "pause"},
		Transition{From: "pause", Event: "logout", To: "login", Action: "logout"},
	)

	assert.True(t, tm.currentState.From != "login")
	assert.False(t, fsm.Login)
	fmt.Println(tm.currentState)

	assert.Nil(t, tm.Process("login"))
	assert.True(t, tm.currentState.From == "run")
	assert.True(t, fsm.Login)
	fmt.Println(tm.currentState)

	assert.Nil(t, tm.Process("pause"))
	assert.True(t, tm.currentState.From != "pause")
	assert.True(t, fsm.Login)
	fmt.Println(tm.currentState)

	assert.Nil(t, tm.Process("login"))
	assert.True(t, tm.currentState.From == "run")
	assert.True(t, fsm.Login)
	fmt.Println(tm.currentState)
}
