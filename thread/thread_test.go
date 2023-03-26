package thread

import (
	"testing"
)

// 测试交替打印字符串
//func TestAlternatePrint(t *testing.T) {
//	res := alternatePrint("12345", "abcde")
//	fmt.Println("测试结果: ", res)
//}

// 测试生产者与消费者
func TestProducerAndConsumer(t *testing.T) {
	producerAndConsumer(5, 5, 5)
}
