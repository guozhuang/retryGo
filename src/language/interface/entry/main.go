package main

import (
	"fmt"
	"language/interface/mock"
	"language/interface/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("http://www.baidu.com")
}

//用来理解interface的duck typing的定义和实现以及interface类型判断
func main() {
	var r Retriever

	r = mock.Retrieve{"this is test mock"}
	//fmt.Println(download(r))
	// 将被调用者的结构体进行调用者接口化处理为接口变量
	// 接口变量进行传参，在调用者函数上声明传入的参数声明调用者的interface参数
	//这样就将被调用者实现了调用者声明的interface
	//从而调用者可以灵活使用被调用者的代码块逻辑【实现了面向接口的调用逻辑】
	fmt.Printf("%T, %v\n", r, r)
	inspect(r)

	r = &real.Retrieve{
		Ua:      "ios",
		TimeOUt: time.Minute,
	}
	fmt.Printf("%T, %v\n", r, r)
	inspect(r)
	//fmt.Println(download(r))

}

//对interface根据类型处理
func inspect(r Retriever) {
	switch v := r.(type) {
	case mock.Retrieve:
		fmt.Println(v.Contents)
	case *real.Retrieve:
		fmt.Println(v.Ua)
	}
}
