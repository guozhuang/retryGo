package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	//实现一个斐波那契数列，并且通过writer接口进行文件写入
	f := feb() //闭包的调用，返回一个带引用的函数（保持引用形成函数内的迭代，此处的引用就是a，b变量）

	/*fmt.Println(f())//对每次进行调用细化处理
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())*/

	//使用了生成器之后就不需要上面那种一直运行的方式，只要一直scan进行循环就能一直调用，并且传入函数
	//生成器自动也进行next函数接收和执行（next = g()）

	//所以使用生成器之后进行函数的触发
	printFileContents(f) //当前就实现了，所以这里的printFileContents函数就能通用化
}

//闭包实现斐波那契数列
func feb() intGen {
	a, b := 1, 0
	return func() int {
		b, a = a, a+b

		return a
	}
}

//创建一个类型来定义返回int（作为一个函数的类型）【可以理解成为生成器】
type intGen func() int

//创建了类型之后，就能进行接口的实现【实现read方法，也就实现了Reader接口】
func (g intGen) Read(p []byte) (n int, err error) {
	next := g() //下一个值

	if next > 10000 {
		return 0, io.EOF
	}

	//使用代理实现byte数组到int值的转换
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p) //通过字符串的NewReader来实现read代理(这里实现就将byte数组操作为标准的read输出)，直接返回该函数的返回值即可
}

//实现一个对reader的读取公共方法，接口reader的接口变量，所以上面斐波那契函数的intGen函数类型实现来reader接口，便能调用该方法
//需要进行read方式读取内容时，直接调用print方法并且传递了实现了接口变量即可
func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
