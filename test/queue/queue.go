package queue

type Queue []int

//slice在函数中传递的时候希望修改值就是下面的方式处理的
//所以在这里的方法中接收者本身是按引用，而且append对应的指针引用
func (q *Queue) Push(v int) {
	//因为针对的是slice的操作，所以需要使用的q的指针进行操作
	*q = append(*q, v) //直接使用接收者的指针来进行指针的修改
}

func (q *Queue) Pop() int {
	head := (*q)[0]

	*q = (*q)[1:]
	return head
}

//接收者设置了指针，内部的slice操作的方法内的变量也都使用了变量的引用
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

/**
slice在函数中进行传递地址，并且操作值
func main(){
	test := []int{1,2,3}

	optSlice(&test)
	for _, v := range test {
		fmt.Println(v)
	}
}

func optSlice(s *[]int){
	*s = append(*s, 4)
}
*/
