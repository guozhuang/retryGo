package queue

type Queue []int

//testing method ptr
func (q *Queue) Push(v int) {
	//q = append(q, v)//注意这里的接收者使用的是指针，内部使用指针的形式进行处理,保证修改的内容将指针对应的值进行了处理
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	first := (*q)[0] //这种写法

	*q = (*q)[1:]

	return first
}

func (q *Queue) IsEmpty() bool {
	if len(*q) == 0 {
		return true
	}

	return false
}
