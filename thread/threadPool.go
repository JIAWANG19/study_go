package thread

import (
	"fmt"
	"sync"
)

// Task 定义任务结构体，包含任务函数
type Task struct {
	f func() error
}

// 定义工作线程结构体
type Worker struct {
	id         int
	taskQueue  chan Task
	workerPool chan chan Task
	quit       chan bool
}

type Pool struct {
	taskQueue   chan Task
	workerPool  chan chan Task
	workers     []*Worker
	maxSize     int
	currentSize int
	wg          sync.WaitGroup
}

func NewWorker(id int, workerPool chan chan Task) Worker {
	return Worker{
		id:         id,
		taskQueue:  make(chan Task),
		workerPool: workerPool,
		quit:       make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.workerPool <- w.taskQueue
			select {
			case task := <-w.taskQueue:
				if err := task.f(); err != nil {
					fmt.Printf("Worker %d: %s\n", w.id, err.Error())
				}
			case <-w.quit:
				fmt.Printf("Worker %d stopping...\n", w.id)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

func NewPool(size int) *Pool {
	pool := &Pool{
		taskQueue:   make(chan Task),
		workerPool:  make(chan chan Task, size),
		maxSize:     size,
		currentSize: 0,
	}
	return pool
}

func (p *Pool) Run() {
	for i := 0; i < p.maxSize; i++ {
		worker := NewWorker(i+1, p.workerPool)
		p.workers = append(p.workers, &worker)
		worker.Start()
	}

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		for {
			select {
			case task := <-p.taskQueue:
				go func() {
					taskQueue := <-p.workerPool
					taskQueue <- task
				}()
			}
		}
	}()
}

func (p *Pool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
	close(p.taskQueue)
	p.wg.Wait()
}

func (p *Pool) AddTTask(f func() error) {
	task := Task{}
	p.taskQueue <- task
}
