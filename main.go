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

	//for i := 0; i < 10; i++ {
	//	pool.AddJob(&ExampleJob{id: i})
	//}
	var i int
	for {
		i++
		pool.AddJob(&ExampleJob{id: i})
		time.Sleep(2 * time.Second)
		if i == 5 {
			pool.Stop()
		}
	}
	pool.Wait()
	pool.Stop()
}
