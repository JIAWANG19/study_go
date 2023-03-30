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
