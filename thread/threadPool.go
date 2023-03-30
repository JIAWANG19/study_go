package thread

import (
	"fmt"
	"sync"
)

type Job interface {
	Do() error
}

type Worker struct {
	id         int
	jobChannel chan Job
	quit       chan bool
}

func NewWorker(id int, jobChannel chan Job) *Worker {
	return &Worker{
		id:         id,
		jobChannel: jobChannel,
		quit:       make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case job := <-w.jobChannel:
				fmt.Printf("Worker %d started job\n", w.id)
				err := job.Do()
				if err != nil {
					fmt.Printf("job has error")
					return
				}
				fmt.Printf("Worker %d finished job\n", w.id)
			case <-w.quit:
				return
			}

		}
	}()
}

func (w *Worker) Stop() {
	w.quit <- true
}

type Pool struct {
	jobChannel chan Job
	workers    []*Worker
	wg         sync.WaitGroup
}

func NewPool(numWorkers int, jobQueueLength int) *Pool {
	jobChannel := make(chan Job, jobQueueLength)
	workers := make([]*Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i, jobChannel)
	}

	return &Pool{
		jobChannel: jobChannel,
		workers:    workers,
	}
}

func (p *Pool) Start() {
	p.wg.Add(1)
	for _, worker := range p.workers {
		worker.Start()
	}
}

func (p *Pool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) AddJob(job Job) {
	p.wg.Add(1)
	p.jobChannel <- job
	p.wg.Done()
}
