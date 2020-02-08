package main

import "fmt"

var G = 7

func main() {

	//首先是针对全局变量的操作【实际使用中应当尽量回避】
	y := func() int {
		fmt.Printf("G:%d, G的地址:%p\n", G, &G)

		G += 1
		return G
	}

	fmt.Println(y(), y)
	fmt.Println(y(), y)
	fmt.Println(y(), y)

	/**
	此时的输出结果为
	G:7, G的地址:0x1167280
	8 0x109b2c0
	G:8, G的地址:0x1167280
	9 0x109b2c0
	G:9, G的地址:0x1167280
	10 0x109b2c0
	*/

	//操作全局变量的匿名函数自调用实现
	z := func() int {
		G += 1
		return G
	}()

	fmt.Println(z, &z)
	fmt.Println(z, &z)

	//此时操作全局变量的自调用可以看出的输出结果始终是一样的
	/**
	11 0xc00001a0d8
	11 0xc00001a0d8
	*/
	//显然，因为此时的z只是执行了一次，已经是直接的输出结果

	//接着就是闭包实现的处理
	var f = N()
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)
	fmt.Println(f(1), &f)

	/**此处的输出结果为：
	i：0, i的地址:0xc00001a0f0
	1 0xc00000e030
	i：1, i的地址:0xc00001a0f0
	2 0xc00000e030
	i：2, i的地址:0xc00001a0f0
	3 0xc00000e030
	*/

	//可以看出，此时的闭包调用每次都是新的执行，只不过闭包的返回函数地址始终是同一个，说明调用的函数也是一致。
	//此时的i变量【在N函数内】能够相对保有一个全局引用的变量

	//而重新调用闭包外层函数时
	var f1 = N()
	fmt.Println(f1(1), &f1)
	/**再次重新生成一个返回的函数地址，并且计数的i互相不影响【也是当然】
	i：0, i的地址:0xc000092098
	1 0xc00008e028
	*/

}

//闭包声明
func N() func(int) int {
	var i int
	return func(d int) int {
		fmt.Printf("i：%d, i的地址:%p\n", i, &i)

		i += d
		return i
	}
}
