package main

import "fmt"

func main() {

	imap := IntMap{}

	imap.LoadOrStore(1, 2)

	if value, ok := imap.Load(1); ok {
		fmt.Println("load")
		fmt.Println(value)
	}
	if value, ok := imap.LoadAndDelete(1); ok {
		fmt.Println("load and delete")
		fmt.Println(value)
	}

}
