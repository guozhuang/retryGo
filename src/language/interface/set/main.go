package main

import "fmt"

type T struct {
	Name string
}

//接口变量的检查方法集和调用情况
type Intf interface {
	M1()
	M2()
}

func (t T) M1() {
	fmt.Println(t.Name)
	t.Name = "name1"
	fmt.Println(t.Name)
}

//使用指针方法
func (t *T) M2() {
	fmt.Println(t.Name)
	t.Name = "name2"

	fmt.Println(t.Name)
}

func main() {
	var t1 T = T{"t1"}
	//t1.M1()
	//t1.M2()//这都是实例的简单方法调用，所以使用指针方法和值方法没有什么差别
	//区别就是指针方法能够进行值的修改
	//t1.M1()//此时输出的实例的name就变为了name2

	//而使用接口变量的话,不使用&t1,则无法编译
	var t2 Intf = &t1
	t2.M1()
	t2.M2()

}
