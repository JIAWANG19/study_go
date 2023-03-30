package main

import (
	"fmt"
	"study_go/thread"
	"time"
)

type ExampleJob struct {
	id int
}

func (ej *ExampleJob) Do() error {
	fmt.Printf("ExampleJob %d started\n", ej.id)
	time.Sleep(time.Second)
	fmt.Printf("ExampleJob %d finished\n", ej.id)
	return nil
}

func main() {
	pool := thread.NewPool(5, 10)
	pool.Start()
	fmt.Println("等待线程池停止指令")
	time.Sleep(2 * time.Second)
	pool.Stop()
}
