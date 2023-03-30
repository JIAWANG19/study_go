package thread

import (
	"fmt"
	"log"
	"sync"
)

// Task 结构体表示一个任务，包含参数 params 和执行函数 f
type Task struct {
	f func() (any, error) // 执行函数
}

// TaskWorker 结构体表示一个任务执行者，包含 ID、任务通道、返回值通道和退出通道
type TaskWorker struct {
	id          int       // 任务执行者的ID
	taskChannel chan Task // 任务通道
	returnValue chan any  // 返回值通道
	quit        chan bool // 退出通道
}

// NewWorkerTask 函数创建一个新的任务执行者
func NewWorkerTask(id int) *TaskWorker {
	return &TaskWorker{
		id:          id,
		taskChannel: make(chan Task), // 创建一个无缓冲通道，用于接收任务
		returnValue: make(chan any),  // 创建一个无缓冲通道，用于返回执行结果
	}
}

// Start 方法启动一个任务执行者，监听任务通道和退出通道，执行接收到的任务并返回结果
func (w *TaskWorker) Start() {
	go func() {
		for {
			select {
			case task := <-w.taskChannel:
				log.Printf("[thread %d]: is working", w.id)
				returnValue, err := task.f() // 执行任务
				w.returnValue <- returnValue // 返回执行结果
				if err != nil {
					log.Printf("[thread %d] task has error", w.id)
				}
				log.Printf("[thread %d]: is ready", w.id)
			case <-w.quit:
				return // 接收到退出通道信号，结束任务执行者
			}
		}
	}()
}

// AddTask 方法向任务通道添加一个新的任务
func (w *TaskWorker) AddTask(params map[string]any, f func() (any, error)) error {
	if len(w.taskChannel) == cap(w.taskChannel) {
		return fmt.Errorf("taskChannel is full") // 任务通道已满，无法添加新的任务
	}
	task := Task{f: f}
	w.taskChannel <- task // 将任务加入到任务通道
	return nil
}

// Stop 方法向退出通道发送信号，停止任务执行者
func (w *TaskWorker) Stop() {
	w.quit <- true // 发送退出信号
}

// GetReturnValue 方法从返回值通道中获取任务执行结果
func (w *TaskWorker) GetReturnValue() any {
	return <-w.returnValue
}

type TaskPool struct {
	workers []*TaskWorker
	wg      sync.WaitGroup
	mu      sync.Mutex
}

func NewTaskPool(numWorkers int) *TaskPool {
	pool := &TaskPool{
		workers: make([]*TaskWorker, numWorkers),
	}
	for i := 0; i < numWorkers; i++ {
		pool.workers[i] = NewWorkerTask(i)
	}
	return pool
}

func (p *TaskPool) Start() {
	p.wg.Add(1)
	for _, worker := range p.workers {
		worker.Start()
	}
}

func (p *TaskPool) Stop() {
	for _, worker := range p.workers {
		worker.Stop()
	}
	p.wg.Done()
}

func (p *TaskPool) Wait() {
	p.wg.Wait()
}

func (p *TaskPool) GetWorker() (*TaskWorker, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	for _, worker := range p.workers {
		if len(worker.taskChannel) == 0 {
			return worker, nil
		}
	}
	return nil, fmt.Errorf("all worker is working")
}
