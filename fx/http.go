package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"go.uber.org/fx"
)

// Logger 構造函數
func NewLogger() *log.Logger {
	logger := log.New(os.Stdout, "" /* prefix */, 0 /* flags*/)
	logger.Print("Executing NewLogger")
	return logger
}

// http.Handler 構造函數
func NewHandler(logger *log.Logger) (http.Handler, error) {
	logger.Print("Executing NewHandler")
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		logger.Print("Got a request.")
	}), nil
}

// http.ServeMux 構造函數
func NewMux(lc fx.Lifecycle, logger *log.Logger) *http.ServeMux {
	logger.Print("Executing NewMux")

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	lc.Append(fx.Hook{ // 自定義生命週期過程對應的啟動和關閉行為
		OnStart: func(ctx context.Context) error {
			logger.Print("Starting HTTP server.")
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Print("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return mux
}

// 註冊 http.Handler
func Register(mux *http.ServeMux, h http.Handler) {
	mux.Handle("/", h)
}

func main() {
	app := fx.New(
		fx.Provide(
			NewLogger,
			NewHandler,
			NewMux,
		),
		fx.Invoke(Register), // 通過 invoke 來完成 Logger、Handler、ServeMux 的創建
	)

	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil { // 手動調用 Start
		log.Fatal(err)
	}

	http.Get("http://localhost:8080/") // 具體操作

	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil { // 手動調用 Stop
		log.Fatal(err)
	}
}
