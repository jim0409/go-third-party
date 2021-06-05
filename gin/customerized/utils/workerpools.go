package utils

import (
	"fmt"
	"log"
)

var Wspools WorkerPools

type WorkRequest struct {
	Execute      func(config interface{}) (string, error)
	ExecuteInput interface{}
	Result       chan interface{}
}

type Worker struct {
	ID         int
	Work       chan WorkRequest
	WorkerPool chan chan WorkRequest
	QuitChan   chan bool
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.Work

			select {
			case work := <-w.Work:
				result, err := work.Execute(work.ExecuteInput)
				if err != nil {
					// log.Printf("Recive error %v", err)
					work.Result <- err
				} else {
					// log.Printf("Processing %v", result)
					work.Result <- result
				}
			case <-w.QuitChan:
				fmt.Printf("worker_%d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

type WorkerPoolType chan chan WorkRequest

func NewDispatcherPool(n int, t string) WorkerPools {
	return &DispatcherPool{
		tag:        t,
		nworkers:   n,
		workerpool: make(WorkerPoolType, n),
		workers:    []*Worker{},
		workqueue:  make(chan WorkRequest),
	}
}

type DispatcherPool struct {
	tag        string
	nworkers   int
	workerpool WorkerPoolType
	workers    []*Worker
	workqueue  chan WorkRequest
}

type WorkerPools interface {
	StartDispatcher()
	StopDispatcher()
	RetriveWorkerQueue() chan WorkRequest
	RetriveWorker() []*Worker
}

func (d *DispatcherPool) StopDispatcher() {
	log.Println(d.tag, " worker pools is going down")
	for _, j := range d.RetriveWorker() {
		j.Stop()
	}
	// time.Sleep(3 * time.Second)
}

func (d *DispatcherPool) StartDispatcher() {
	log.Println(d.tag, " worker pools is up")

	for i := 0; i < d.nworkers; i++ {
		fmt.Println("Starting worker ", i+1)
		worker := d.NewWorker(i)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-d.workqueue:
				// log.Println("Received work ...")
				go func() {
					worker := <-d.workerpool

					// log.Println("Dispatching work request")
					worker <- work
				}()
			}

		}
	}()

}

// 已經封裝workerpool，所以不需要在帶入輸入參數
// func (d *DispatcherPool) NewWorker(id int, workerQueue chan chan WorkRequest) *Worker {
func (d *DispatcherPool) NewWorker(id int) *Worker {
	worker := &Worker{
		ID:         id,
		Work:       make(chan WorkRequest),
		WorkerPool: d.workerpool,
		QuitChan:   make(chan bool),
	}
	d.workers = append(d.workers, worker)

	return worker
}

func (d *DispatcherPool) RetriveWorkerQueue() chan WorkRequest {
	return d.workqueue
}

// RetriveWorker would return an array with Workers
func (d *DispatcherPool) RetriveWorker() []*Worker {
	return d.workers
}
