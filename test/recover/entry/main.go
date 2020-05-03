package main

import (
	"fmt"
)

func main() {
	tryRecover()
}

func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			//说明此处捕获的panic的值是一个error类型
			fmt.Println("error occurred:", err)
		} else {
			panic(r)
		}
	}()

	//panic(errors.New("this is err"))

	/*b := 0
	a := 5 / b
	fmt.Println(a)//也正常recover，不会出现panic
	*/
	panic(123) //直接panic，并且在defer中因为recover不到，导致接着panic
	/**
	panic: 123 [recovered]
	panic: 123
	*/
}
