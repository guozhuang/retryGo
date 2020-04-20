package test

//实现类似用途函数的封装和调用灵活【apply】
func Add(a, b int) int {
	return a + b
}

func Min(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

/**
再后面使用接口来约束相应的方法，并且进行接口方法的代理实现，同样使用的也是函数式编程的思路
【更加灵活地约束和接口实现】
result := test.Apply(test.Mul, 3, 5)

	fmt.Println(result)
*/
