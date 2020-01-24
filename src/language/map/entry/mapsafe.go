package main

import (
	"fmt"
	"strconv"
	"sync"
)

//map 线程安全实现

func main() {
	wg := sync.WaitGroup{}
	//使用协程来对map进行处理
	//使用sync.map
	//c := make(map[string]string)
	var c sync.Map

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			k, v := strconv.Itoa(n), strconv.Itoa(n)
			c.Store(k, v)
			//
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(c)

	//进行该结构的读取
	for i := 0; i < 10; i++ {
		if data, ok := c.Load(strconv.Itoa(i)); ok {
			fmt.Println("key", data)
		}
	}
}
