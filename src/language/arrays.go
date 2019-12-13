package main

import "fmt"

//这里的入参不声明长度的话，会表示报错
func optMultiArr(args [4][5]int) {
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			fmt.Println(args[i][j])
		}
	}
}

func main() {
	multiArr := [4][5]int{}

	optMultiArr(multiArr)
	//
}
