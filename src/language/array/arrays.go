package array

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
	/*multiArr := [4][5]int{}

	optMultiArr(multiArr)*/

	box := [256]int{} //创建定长数组
	for k, _ := range box {
		box[k] = k
	}
	fmt.Println(box)
}
