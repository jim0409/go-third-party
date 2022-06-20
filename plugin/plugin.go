//go:build plugin
// +build plugin

package main

import (
	"fmt"
	"go-third-party/plugin/plugin"
)

func init() {
	fmt.Println("load plugin.so success!!!")
}

func main() {
	plugin.Run()
}

func Run() {
	plugin.Run()
}
