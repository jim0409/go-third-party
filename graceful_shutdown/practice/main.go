package main

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/appleboy/graceful"
)

func main() {
	m := graceful.NewManager()

	ln, _ := net.Listen("tcp", ":8200")
	http.HandleFunc("/count", count)

	/*
		用 http server 會卡住
		因為 net/http 底層會有聽 signal interrupt 事件了
	*/
	/*
		m.AddRunningJob(func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return http.ErrServerClosed
			default:
				return http.Serve(ln, nil)
			}
		})
	*/

	go http.Serve(ln, nil)

	/*
		很適合用來做副執行緒
	*/
	m.AddRunningJob(print)

	<-m.Done()
}

func count(res http.ResponseWriter, r *http.Request) {
	// log.Println("count++")
	atomic.AddInt32(&num, 1)
	_, err := io.WriteString(res, "ok")
	if err != nil {
		log.Fatal(err)
	}
}

var num int32

func print(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			log.Printf("now # %d\n", num)
			time.Sleep(time.Second)
		}
	}
}
