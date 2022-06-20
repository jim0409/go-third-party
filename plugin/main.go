//go:build release
// +build release

package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("./plugin.so")
	if err != nil {
		fmt.Printf("load plugin failed err:%v\n", err)
		return
	}

	runFunc, err := p.Lookup("Run")
	if err != nil {
		fmt.Printf("lookup Run func err:%v\n", err)
		return
	}

	runFunc.(func())()
}
