package main

import (
	"fmt"
	"language/queue"
)

//testing queue package
//look up package and method
func main() {
	q := queue.Queue{}

	q.Push(1)
	q.Push(100)

	fmt.Println(q.Pop())

	fmt.Println(q.IsEmpty())

	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
