package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fibomacci() intGen {
	a, b := 0, 1

	return func() int {
		a, b = b, a+b //golang的交换变量更简单
		return a
	}
}

//将斐波那契处理函数封装成一个文件读的函数，将其封装在内，并且对外提供函数接口来调用
//将闭包函数调用之后的结果，直接传递到对应函数类型（接口）上，进而实现函数的通用打印【打印的方法即实现了定义的函数类型的接口】

//定义一个函数类型
type intGen func() int

//此处声明一个实现函数类型的方法
//所以调用io.Reader接口的时候，因为intGen函数类型已经具备了Read方法，
// 所以符合intGen的函数就符合下面printFileContents函数中的参数中要求的接口类型
//所以斐波那契的闭包函数能够传递给printFileContents函数【并且代理了写入文件，然后再通用read该函数的处理结果】
func (g intGen) Read(p []byte) (n int, err error) {
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	return strings.NewReader(s).Read(p) //注意这里的方法
}

//此处的io.Reader接口仅需要实现Read方法便实现了该接口
func printFileContents(r io.Reader) {
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

//重点：为函数实现接口
func main() {
	f := fibomacci()

	/*fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/

	printFileContents(f)
}
