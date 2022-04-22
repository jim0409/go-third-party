package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Consumer struct
type Consumer struct {
	inputChan chan int
	jobsChan  chan int
}

func getRandomTime() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(10)
}

// withContextFunc: 使用 context 做文本管控，如果文本狀態為結束則不做任何執行，反之執行取消文本及`f`函數做結束
func withContextFunc(ctx context.Context, f func()) context.Context {
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(c)

		select {
		case <-ctx.Done():
		case <-c:
			cancel()
			f()
		}
	}()

	return ctx
}

// queue: 將 input 丟進 Consumer 的 inputChan
func (c *Consumer) queue(input int) bool {
	select {
	case c.inputChan <- input:
		log.Println("already send input value:", input)
		return true
	default:
		return false
	}
}

// startConsumer: 告知所有 consumer 開始進行 inputChan 的消費
// 如果收到 文本的 ctx.Done 則關閉該 Consumer 的 jobsChan
func (c Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.inputChan:
			if ctx.Err() != nil {
				close(c.jobsChan)
				return
			}
			c.jobsChan <- job
		case <-ctx.Done():
			close(c.jobsChan)
			return
		}
	}
}

// process: 每個 consumer 都會執行的工作，模擬實際運行時執行的事情
func (c *Consumer) process(num, job int) {
	n := getRandomTime()
	log.Printf("Sleeping %d seconds...\n", n)
	time.Sleep(time.Duration(n) * time.Second)
	log.Println("worker:", num, " job value:", job)
}

func (c *Consumer) worker(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Println("start the worker", num)
	for {
		select {
		case job := <-c.jobsChan:
			if ctx.Err() != nil {
				log.Println("get next job", job, "and close the worker", num)
				return
			}
			c.process(num, job)
		case <-ctx.Done():
			log.Println("close the worker", num)
			return
		}
	}
}

const poolSize = 2

func main() {
	finished := make(chan bool)
	wg := &sync.WaitGroup{}
	wg.Add(poolSize)
	// create the consumer
	consumer := Consumer{
		inputChan: make(chan int, 10),
		jobsChan:  make(chan int, poolSize),
	}

	ctx := withContextFunc(context.Background(), func() {
		log.Println("cancel from ctrl+c event")
		wg.Wait()
		close(finished)
	})

	for i := 0; i < poolSize; i++ {
		go consumer.worker(ctx, i, wg)
	}

	go consumer.startConsumer(ctx)

	go func() {
		consumer.queue(1)
		consumer.queue(2)
		consumer.queue(3)
		consumer.queue(4)
		consumer.queue(5)
	}()

	<-finished
	log.Println("Game over")
}
