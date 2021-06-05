package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zommage/cron"
)

const (
	// 秒 分 時 日 月
	cron1 = "*/1 * * * * ?"
)

func main() {
	cronRes := cron.New()
	if err := addCrons(cronRes); err != nil {
		panic(err)
	}
	cronRes.Start()
	defer cronRes.Stop()
	gracefulShutdown()
}

func addCrons(c *cron.Cron) error {
	c.AddFunc(cron1, func() {
		go printInt()
	})
	return nil
}

func printInt() {
	fmt.Println(time.Now())
}

func gracefulShutdown() {

	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)
	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
