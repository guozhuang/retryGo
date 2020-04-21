package main

import (
	"fmt"
	"retryGo/test/queue"
)

func main() {
	q := queue.Queue{}
	q.Push(1)
	q.Push(2)
	q.Push(2)
	q.Push(3)

	for !q.IsEmpty() {
		data := q.Pop()
		fmt.Println(data)
	}
}
