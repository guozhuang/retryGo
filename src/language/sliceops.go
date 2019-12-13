package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("this is %v, len: %d; cap: %d \n", s, len(s), cap(s))
}

func pop(s []int) []int {
	var re []int
	if len(s) > 0 {
		re = s[:len(s)-1]
	}

	return re
}

func shift(s []int) []int {
	var re []int

	if len(s) > 0 {
		re = s[1:]
	}

	return re
}

func main() {
	//slice的自动扩容，还是两倍的递增处理，这一点还是为了优化扩容的次数
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{1, 2}

	s4 := []int{3, 4}

	//使用make关键字来创建
	s2 := make([]int, 10)
	printSlice(s2) //[0 0 0 0 0 0 0 0 0 0], len: 10; cap: 10

	s3 := make([]int, 5, 20)
	printSlice(s3) //[0 0 0 0 0], len: 5; cap: 20

	copy(s2, s1)
	printSlice(s2) //[1 2 0 0 0 0 0 0 0 0], len: 10; cap: 10

	copy(s2, s4)   //始终是追加到头部
	printSlice(s2) //[3 4 0 0 0 0 0 0 0 0], len: 10; cap: 10

	copy(s4, s1)   //这里会进行替换
	printSlice(s4) //[1 2], len: 2; cap: 2

	sQueue := []int{1, 2, 3, 4, 5}
	sPop := pop(sQueue) //后面再追加被操作的元素作为返回值
	fmt.Println("pop data", sPop)

	sShift := shift(sQueue)
	fmt.Println("shift data", sShift)
}
