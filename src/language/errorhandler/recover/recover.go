package main

import (
	"fmt"
)

func tryRecover() {
	defer func() {
		r := recover()

		if err, ok := r.(error); ok {
			fmt.Println("get panic error: ", err)
		} else {
			panic(r) //无法recover,接着panic
		}
	}()

	//panic(errors.New("this is opt panic"))//panic传入err，然后在recover中获取

	//panic(123)//无法探知

	b := 0
	a := 5 / b
	fmt.Println(a) //触发了系统默认的panic报err的情况，走到了recover上，没有直接挂掉程序
}

func main() {
	tryRecover()
}
