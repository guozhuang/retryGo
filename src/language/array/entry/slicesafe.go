package main

import (
	"fmt"
	"sync"
)

//slice 的线程安全问题

func main() {
	num := 10000

	var wg sync.WaitGroup
	wg.Add(num)

	c := make(chan int)
	for i := 0; i < num; i++ {
		go func() {
			c <- 1 // channl是协程安全的
			wg.Done()
		}()
	}

	// 等待关闭channel
	go func() {
		wg.Wait()
		close(c)
	}()

	// 读取数据
	var a []int
	for i := range c {
		a = append(a, i)
	}

	fmt.Println(len(a))
}
