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

func main() {
	//catch global panic
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic err: ", err)
		}
	}()

	cs := newChatServer()
	s := &http.Server{
		Addr:    ":8080",
		Handler: cs,
		// ReadTimeout:  time.Second * 10,
		// WriteTimeout: time.Second * 10,
	}
	errc := make(chan error, 1)
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("http listen : %v\n", err)
		}
	}()

	cs.cronJob()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
