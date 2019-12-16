package apply

import "fmt"

func addNum(a int, b int) int {
	var sum int
	sum = a + b
	return sum
}

func minus(a, b int) int {
	return a - b
}

//使用这种形式来实现函数解耦,进而对照相应的函数式传参，实现对应的接口，来实现进一步的解耦
func apply(op func(int, int) int, a, b int) int {
	//函数作为参数就是函数式编程的重要实现，从而实现的解耦。。。正是因为有了函数式编程，面向接口变得比较容易实现
	return op(a, b) //这里直接调用传入的函数，也可以在函数内使用闭包的定义来实现
}

func opt(opt string, a, b int) (int, error) {
	switch opt {
	case "+":
		return apply(addNum, a, b), nil
	case "-":
		return apply(minus, a, b), nil
	default:
		return 0, fmt.Errorf("err opt")
	}
}

func main() {
	sum, _ := opt("+", 3, 1)
	fmt.Println(sum)

	min, _ := opt("-", 3, 1)
	fmt.Println(min)

	fmt.Println(opt("*", 3, 2)) //新增功能的时候只需要新增一个函数，然后调用时再加上即可正常使用
}
