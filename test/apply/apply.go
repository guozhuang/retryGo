package apply

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
扩展新的功能，只需要按照约束的函数参数类型和返回类型之外，新增一个实现了该接口的函数，即可完成了功能的扩展
【这样的灵活性就是函数式编程所提倡的】
result := test.Apply(test.Mul, 3, 5)

	fmt.Println(result)
*/
