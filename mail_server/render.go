package main

import "fmt"

var (
	simpleStringRender = `hello %v`
)

var MsgTemplate = make(map[int]string)

func init() {
	MsgTemplate[0] = fmt.Sprintf(simpleStringRender, "Jim")
}
