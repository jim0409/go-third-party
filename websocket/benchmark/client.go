package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"nhooyr.io/websocket"
)

func main() {

	t1 := time.Now().Unix()

	// var wg *sync.WaitGroup
	wg := sync.WaitGroup{}

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			c, _, err := websocket.Dial(ctx, "ws://127.0.0.1:3000/ws/conn", &websocket.DialOptions{
				Subprotocols: []string{"echo"},
			})
			if err != nil {
				panic(err)
			}
			defer cancel()

			c.Close(websocket.StatusNormalClosure, "")
			defer wg.Done()
		}()
	}
	wg.Wait()

	t2 := time.Now().Unix()

	fmt.Println(t2 - t1)
}
