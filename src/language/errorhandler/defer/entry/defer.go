package main

import "fmt"

func tryDefer() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("enough")
		}
	}
}

func main() {
	tryDefer()
}
