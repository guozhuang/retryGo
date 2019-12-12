package main

import (
	"fmt"
	"strconv"
)

//枚举式
func enums() {
	const (
		m = 1 * (1000 * iota)
		km
	)

	fmt.Println(m, km)
}

//数字转二进制
func converToBin(num int) string {
	var result string

	//初始条件省略：lab:=0
	for ; num > 0; num = num / 2 {
		lab := num % 2

		result = strconv.Itoa(lab) + result
	}
	return result
}

//语言特点记录
func main() {
	enums()

	fmt.Println(
		converToBin(5),
		converToBin(13),
	)
}
