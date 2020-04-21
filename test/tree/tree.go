package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) GetVal() {
	fmt.Println(node.Value)
}

//设计用来pipeline方式调用
func (node *Node) SetVal(val int) *Node {
	node.Value = val
	//因为返回来对应的实例来实现链式调用：node.SetVal(5).GetVal()//这样就实现包内方法的链式调用
	return node
}
