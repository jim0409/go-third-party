package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func serve(ctx context.Context) (err error) {

	// 一般流程，宣告一個server物件
	mux := http.NewServeMux()

	// 一般流程，替該物件增加新的HandlerFunc
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "okay")
		},
	))

	// 用http.Server把之前宣告的物件包裝起來
	srv := &http.Server{
		Addr:    ":6969",
		Handler: mux,
	}

	// 執行server
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("server started")

	// 如果執行cancel()的話，ctx.Done()會解除擁塞
	<-ctx.Done()

	log.Printf("server stopped")

	// 宣告另一個文本，用來做延遲請求處理
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// 延遲處理文本
	defer func() {
		cancel()
	}()

	// 關閉伺服器
	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

func main() {
	// 使用c這個變量來聽取傳遞訊號
	c := make(chan os.Signal, 1)
	// 當程序接收到取消指令時，會賦予c取消的訊號
	signal.Notify(c, os.Interrupt)

	// 在背景，儲存文本。當取消的時候，文本會紀錄 ps : cancel為First-class function
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		// 當發送取消給程序的時候，文本會做取消
		cancel()
	}()

	// 將先前製作的文本傳遞到serve內
	if err := serve(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
