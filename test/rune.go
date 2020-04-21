package test

import (
	"fmt"
	"unicode/utf8"
)

func testRune() {
	s := "hello go语言"
	fmt.Printf("string 长度：%d\n", len(s)) //14
	// hello+空格+go整体占8个字符，所以就是8字节，然后"语言"这两个字符因为使用utf8，每个中文都占3个字节，所以占6个字节。于是总体14字节
	//而单个rune是4个字节【可变，例如英文依旧是一个字节】，能够覆盖所有的utf8编码的字符

	//
	for _, v := range []byte(s) {
		fmt.Println(v)
	}

	fmt.Println()
	for i, v := range s {
		fmt.Printf("%d, %x\n", i, v) //这就能看的出来每个字符占的起始位置
		//此时的v就是一个rune【】
	}

	fmt.Println(utf8.RuneCountInString(s)) //获得的就是字符串的个数

	//使用byte数组转为rune，并且进行展示【这一点在io场景下非常常见】
	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes) //逐个进行decode
		bytes = bytes[size:]               //进行reslice//进行迭代
		fmt.Printf("%c\n", ch)             //实现了字符数组转为可识别的字符内容
	}

	//直接将字符串转化为rune数组，因为rune的count就是对应的字符的数量，所以直接进行rune数组来迭代来处理字符【防止乱码】
	for i, ch := range []rune(s) {
		fmt.Printf("%d, %c\n", i, ch)
	}
}
