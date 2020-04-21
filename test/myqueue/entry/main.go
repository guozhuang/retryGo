package main

import (
	"fmt"
	"retryGo/test/myqueue"
)

func main() {
	my := myqueue.My{Key: "hello", Value: []int{1}}

	my.Push("hello", 2)

	for i := 0; i < 2; i++ {
		fmt.Println(my.Pop("hello"))
	}
}
