package basic

import (
	"fmt"
	"testing"
)

func TestDemo1(t *testing.T) {
	// 可以初始化值
	var map1 = map[string]string{
		"key1": "value1",
		"key2": "value2",
	}

	// 这段代码是有问题的。new 函数返回的是一个指向新分配类型零值的指针，
	// 对于 map 类型的指针来说，其零值是 nil，表示这个指针并没有指向任何有效的 map 实例。
	// 在这段代码中，map2 被初始化为一个 map 指针类型的零值 nil，
	// 然后在其上执行 map2["key1"] = "value1" 操作将导致 panic。
	// var map2 = *new(map[string]string)
	// map2["key1"] = "value1"

	// 可以初始化容量，不可以初始化值
	var map3 = make(map[string]string, 3)
	map3["key1"] = "value1"

	fmt.Println(map1)
	fmt.Println(map3)
}
