package main

import "log"

// Work is something that can be done by workers
type Work interface {
	Execute() error
}

type WorkerPool []Worker

func (pool WorkerPool) StopWorkers() {
	for _, worker := range pool {
		worker.Stop()
	}
}

// Worker can do work
type Worker struct {
	WorkerPool  WorkerPool
	WorkChannel chan Work
	quit        chan bool
}

func NewWorker(pool WorkerPool) Worker {
	return Worker{
		WorkerPool:  pool,
		WorkChannel: make(chan Work),
		quit:        make(chan bool),
	}
}

// Start method starts the run loop for the worker, listening for a quit channel in
// case we need to stop it
func (worker Worker) Start() {
	go func() {
		for {
			// register the current worker in the registry
			worker.WorkerPool = append(worker.WorkerPool, worker)

			select {
			case work := <-worker.WorkChannel:
				err := work.Execute()
				if err != nil {
					log.Print("error", err.Error())
				}
			case <-worker.quit:
				log.Print("stopping worker")
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests
func (worker Worker) Stop() {
	go func() {
		worker.quit <- true
	}()
}
