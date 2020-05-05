package main

import (
	"fmt"
	"time"
)

func worker(id int, ch chan int) {
	for {
		//从channel中读数据
		fmt.Printf("worker %d data is %c\n", id, <-ch)
	}
}

func main() {

	//channel+worker交互
	var channels [10]chan int

	//进行任务的基本分发
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)

		go worker(i, channels[i])
	}

	//进行数据填充
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i //单引号的字符使用
	}

	//重复填充
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i //单引号的字符使用
	}

	time.Sleep(time.Microsecond * 1000)
}
