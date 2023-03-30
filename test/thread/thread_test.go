package thread

import (
	"fmt"
	"study_go/thread"
	"testing"
)

// 测试交替打印字符串
func TestAlternatePrint(t *testing.T) {
	res := thread.AlternatePrint("12345", "abcde")
	fmt.Println("测试结果: ", res)
}

// 测试生产者与消费者
func TestProducerAndConsumer(t *testing.T) {
	thread.ProducerAndConsumer(5, 5, 5)
}

func TestPoolTask(t *testing.T) {
	pool := thread.NewTaskPool(2)
	pool.Start()
	arr := []int{1, 2, 3, 4}
	worker1, _ := pool.GetWorker()
	//worker2, err := pool.GetWorker()
	err := worker1.AddTask(map[string]any{"a": 1}, func() (any, error) {
		fmt.Println(arr)
		return nil, nil
	})
	if err != nil {
		return
	}
	pool.Stop()
}
