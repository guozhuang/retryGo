package main

import "fmt"

func main() {
	test := []int{1, 2, 3}

	optSlice(&test)
	for _, v := range test {
		fmt.Println(v)
	}
}

func optSlice(s *[]int) {
	*s = append(*s, 4)
}
