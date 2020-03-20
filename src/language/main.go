package main

import "fmt"

func main() {
	x := 1
	y := 2

	rex, rey := intswap(x, y)

	fmt.Println(rex)
	fmt.Println(rey)
}

func intswap(x, y int) (int, int) {
	//a^b^a = b
	x ^= y
	y ^= x
	x ^= y
	return x, y
}
