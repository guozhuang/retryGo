package main

import (
	"fmt"
	"retryGo/test/oop/entry/mock"
	"retryGo/test/oop/entry/real"
)

//定义一个接口
type Retriever interface {
	//定义方法集
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

func main() {
	//使用接口变量
	var r Retriever
	//此处将接口变量和被调用者进行绑定【这也就是所谓编译时绑定】
	//mock:因为Get方法接收者是指针接收者，所以这里的接口变量赋值需要使用引用，如果不使用地址引用就直接报错。
	//【所以看起来是指针接收者对应的方法集更为广泛，实际上就是针对这里的接口变量的可调用的范围更广，也就是ide本身也会过滤实现接口的方法集】
	r = &mock.Retrieve{
		Contents: "this is mock result",
	}

	//对接口变量的类型%T可以看出其类型的不同，于是又引出了通过接口来实现多态【type】
	fmt.Printf("%T, %v\n", r, r)

	//real
	r = &real.Retrieve{
		UserAgent: "ua",
		TimeOut:   3,
	} //实际调用正常实现【所以golang中调用方决定被使用者实现了Retriever接口，于是接口变量的赋值来进行调用方】
	fmt.Printf("%T, %v\n", r, r)
	//res := download(r)//此时需要被调用者的实现Get方法，通过接口变量的切换就能切换被调用者处理的流

	//fmt.Println(res)

	//接口变量的多态判断【因为存在接口变量被多个类型进行赋值，而需要对接口变量类型进行判断，便是多个实例的多态判断实现】
	//所以使用接口来实现多态模式
	switch v := r.(type) {
	case *mock.Retrieve: //需要对应这里的指针
		fmt.Println("mock", v.Contents)
	case *real.Retrieve:
		fmt.Println("real", v.UserAgent)
	}
}
