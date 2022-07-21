package main

import (
	"fmt"
	"os"
	"testing"
)

const (
	Logout = "LOGOUT"
	Login  = "LOGIN"
	Start  = "START"
	Play   = "PLAY"
	Pause  = "PAUSE"
	End    = "END"
)

func TestStateMachine(t *testing.T) {
	fm := &Machine{
		Handlers:   make(map[string]Handler),
		StartState: Login,
		EndStates:  make(map[string]bool),
	}

	// ad state
	fm.AddState(Login, LoginHandler)
	fm.AddState(Start, StartHandler)
	fm.AddState(Play, PlayHandler)
	fm.AddState(Pause, PauseHandler)
	fm.AddState(End, EndHandler)
	fm.AddState(Logout, LogoutHandler)

	// add end state
	fm.AddEndState(Logout)

	fm.Execute(Play)
	fmt.Println("====== finished =======")

}

func Cargo() interface{} {
	return LoginHandler
}

func LoginHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute login with cargo status", i)
	return Start, StartHandler
}

func StartHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute start with cargo status", i)
	return Play, PlayHandler
}

var count = 0

func PlayHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute play with cargo status", i, count)
	count = count + 1
	if count < 3 {
		return Play, PlayHandler
	}
	return Pause, PauseHandler
}

func PauseHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute pause with cargo status", i)
	return End, EndHandler
}

func EndHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute end with cargo status", i)
	return Logout, LogoutHandler
}

func LogoutHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute logout with cargo status", i)
	return Logout, nil
}

func EndOneLoop() {
	os.Exit(0)
}
