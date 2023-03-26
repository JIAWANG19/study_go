package thread

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

// 交替打印字符串
func alternatePrint(s1, s2 string) string {
	if len(s1) != len(s2) {
		panic("字符串的长度必须相等，不然怎么交替打印？")
	}
	var num = make(chan int, 1)
	var char = make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(2)
	var res []byte
	go func() {
		defer wg.Done()
		for _, c := range s1 {
			res = append(res, byte(c))
			num <- 1
			<-char
		}
	}()

	go func() {
		defer wg.Done()
		for _, c := range s2 {
			<-num
			res = append(res, byte(c))
			char <- 1
		}
	}()
	wg.Wait()
	return string(res)
}

func producerAndConsumer(producerNum, consumerNum, num int) {
	const letterBytes = "abcdefghijklmnopqrstuvwxyz"
	product := make(chan string, num)
	producer := func(name string) {
		for {
			strLen := 1 + rand.Intn(9)
			bytes := make([]byte, strLen)
			for i := range bytes {
				bytes[i] = letterBytes[rand.Intn(len(letterBytes))]
			}
			productName := string(bytes)
			product <- productName
			log.Printf("[%s] 生产者生产了一个产品 [%s]", name, productName)
			time.Sleep(1 * time.Second)
		}
	}
	consumer := func(name string) {
		for {
			productName := <-product
			log.Printf("[%s] 消费者消费了一个产品 [%s]", name, productName)
			time.Sleep(1 * time.Second)
		}
	}
	for i := 0; i < producerNum; i++ {
		go producer(fmt.Sprintf("producer %d", i))
	}

	for i := 0; i < consumerNum; i++ {
		go consumer(fmt.Sprintf("consumer %d", i))
	}
	time.Sleep(10 * time.Second)
}
