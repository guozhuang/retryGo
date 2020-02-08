package main

import "fmt"

//此处的函数中返回值就包含了具名变量
func func1() (i int) {
	defer func() {
		i = i + 10 //此处的defer就进行的具名返回值的读取和修改，可以看出defer的执行在返回值赋值和return之间进行的
	}()

	return 0
}

func main() {
	fmt.Println("result=>", func1())
}
