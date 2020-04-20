package main

import (
	"fmt"
	"retryGo/test"
)

func main() {
	result := test.Apply(test.Mul, 3, 5)

	fmt.Println(result)
}
