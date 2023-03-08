package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/robfig/cron"
)

const (
	// 秒 分 時 日 月
	cron1 = "*/1 * * * * ?"
	cron2 = "*/2 * * * * ?"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	cronjob := cron.New()
	cronjob.AddFunc(cron1, func() { fmt.Println("test1") })
	cronjob.AddFunc(cron1, func() { fmt.Println("test2") })
	cronjob.AddFunc(cron2, func() { wg.Done() })
	cronjob.Start()
	defer cronjob.Stop()

	select {
	// 期許要在4秒內結束
	case <-time.After(4 * time.Second):
		fmt.Println("expected job fires 2 times")
	// 每2秒完成一個 wait, 4秒完成2個
	case <-wait(wg):
		fmt.Println("finish 2 wait")
	}

}

func wait(wg *sync.WaitGroup) chan bool {
	ch := make(chan bool)
	go func() {
		wg.Wait()
		ch <- true
	}()
	return ch
}
