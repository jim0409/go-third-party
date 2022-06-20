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
	Pause  = "PAUSE"
	End    = "END"
)

func TestStateMachine(t *testing.T) {
	fm := &Machine{
		Handlers:   make(map[string]Handler),
		StartState: Logout,
		EndStates:  make(map[string]bool),
	}

	// ad state
	fm.AddState(Logout, LoginHandler)
	fm.AddState(Login, StartHandler)
	fm.AddState(Start, PauseHandler)
	fm.AddState(Pause, EndHandler)
	fm.AddState(End, LogoutHandler)

	// add end state
	fm.AddEndState(Logout)

	fm.Execute(Cargo)
	fmt.Println("====== finished =======")

}

func Cargo() interface{} {
	return LoginHandler
}

func LoginHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute login")
	return Login, StartHandler
}

func StartHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute start")
	return Start, PauseHandler
}

func PauseHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute pause")
	return Pause, EndHandler
}

func EndHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute end")
	return End, LogoutHandler
}

func LogoutHandler(i interface{}) (string, interface{}) {
	fmt.Println("execute logout")
	return Logout, nil
}

func EndOneLoop() {
	os.Exit(0)
}
