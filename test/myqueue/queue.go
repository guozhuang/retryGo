package myqueue

//构造一个多路的队列
type My struct {
	Key   string
	Value []int //如果此处切换为指针，那么下面的my.Value都需要切换为指针的形式
}

func (my *My) Push(key string, data int) {
	if my.Key == key {
		my.Value = append(my.Value, data)
	}
}

func (my *My) Pop(key string) int {
	if my.Key == key {
		data := (my.Value)[0]
		my.Value = (my.Value)[1:]

		return data
	}

	return 0
}
