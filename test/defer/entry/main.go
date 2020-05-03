package main

import (
	"fmt"
	"os"
)

func main() {
	//deferNormal()
	value := returnValues()
	fmt.Println(value) //此时返回0

	value1 := namedReturnValues()
	fmt.Println(value1) //此时返回1

	deferExit()
}

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer", result)
	}()
	return result
}

func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer", result)
	}()
	return result
}

func deferExit() {
	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}

//测试defer的多个场景已经关联的逻辑
func deferNormal() {
	for i := 0; i < 100; i++ {
		defer fmt.Println("defer num", i) //【因为defer已经声明了30次】这里逐个打印i的数。【表现出的特性就是defer的i值以及defer的先进后出的原则】

		if i == 30 {
			panic("end") //主动进行退出
		}
	}
}
