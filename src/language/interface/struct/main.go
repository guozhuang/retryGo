package main

import "fmt"

type Writer interface {
	Write()
}

type Author struct {
	name   string
	Writer //此处结构体嵌入interface
}

func (author Author) Write() {
	fmt.Println(author.name, "write")
}

type Other struct {
	i int
}

func (other Other) Write() {
	fmt.Println(other.i, "write")
}

//此时两个结构体就能通过interface进行组合操作
func main() {
	//将other结构体的实例通过interface进行处理
	// 【此处就将other结构体变量赋值给对应的接口Writer上，编译器本身会进行检查，other结构体变量是否实现了Writer接口】
	au := Author{"other", Other{99}}
	au.Write() //other write
}
