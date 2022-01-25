package main

import "github.com/rs/xid"

func main() {
	guid := xid.New()

	println(guid.String())
}
