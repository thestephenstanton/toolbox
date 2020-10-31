package main

type Dispatcher struct {
	maxWorkers int
	workerPool chan Worker
	workQueue  chan Work
}

// func NewDispatcher(maxWorkers int) *Dispatcher {
// 	workerPool := make(WorkerPool, 0, maxWorkers)

// 	return &Dispatcher{
// 		maxWorkers: maxWorkers,
// 		workerPool: workerPool,
// 	}
// }

// func (dispatcher *Dispatcher) Run() {
// 	for i := 0; i < dispatcher.maxWorkers; i++ {
// 		worker := NewWorker(dispatcher.workerPool)

// 		worker.Start()
// 	}

// 	go dispatcher.dispatch()
// }

// func (dispatcher *Dispatcher) dispatch() {
// 	for {
// 		select {
// 		case work := <-dispatcher.WorkQueue:
// 			go func(work Work) {
// 				// try to obtain a worker job channel that is available.
// 				// this will block until a worker is idle
// 				worker := <-dispatcher.workerPool

// 				// dispatch the work to the worker job channel
// 				jobChannel <- job
// 			}(work)
// 		}
// 	}
// }
