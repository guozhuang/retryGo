package main

import "fmt"

func main() {
	i := 1

	defer fmt.Println("result=>", func() int { return i * 2 }()) //这里就使用了defer的规则一，此时的i变量的值为1

	i++

	defer func() {
		fmt.Println("result1=>", i*2) //此处对照的defer的规则二，所以先输出当前的值，并且处于两次++操作之后的执行
	}()

	i++
}
