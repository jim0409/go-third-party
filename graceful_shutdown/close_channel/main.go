package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func() {
		ch <- 1
		ch <- 2
		close(ch) // 沒有關閉 channel 會導致 deadlock
	}()

	for n := range ch {
		fmt.Println(n)
	}
}
