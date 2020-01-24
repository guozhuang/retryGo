package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//字符转换处理
	s := "momo科技"

	//普通的byte处理
	for i, ch := range []byte(s) {
		fmt.Println("string data ", i, ch)
	}
	fmt.Println(len(s))

	for i, ch := range s {
		//此时的ch就是一个rune类型的字符描述
		fmt.Println("rune data", i, ch)
		//这里的i描述的下标是根据字符存储的utf8来占位的
	}

	//获取字符占位长度
	fmt.Println("rune count", utf8.RuneCount([]byte(s))) //获取的就是字符个数，汉字也是1位

	//进行每个字符的获取
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //匹配到一个字符之后进行返回，并且返回该字符的长度
		bytes = bytes[size:]               //表示逐个字符的处理
		fmt.Printf("%c\n", ch)
	}

	//直接使用rune来获取每个字符的处理【将上面的方式变得更通用】【这样就比较好理解rune类型来描述字符】
	for i, ch := range []rune(s) {
		fmt.Printf("%d %c \n", i, ch)
	}
}
