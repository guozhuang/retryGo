package main

import "fmt"

func setName(name *string) {
	*name = "world"
}

func addPtr(num *int) {
	*num = *num + 1
}

func notPtr(num *int) {
	//num = num + 1
	//这才是指针不能进行运算的形式
}
func main() {
	name := "hello"
	setName(&name)
	fmt.Println(name)

	a := 1
	addPtr(&a)
	fmt.Println(a)
}
