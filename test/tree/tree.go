package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func (node *TreeNode) GetVal() {
	fmt.Println(node.Value)
}

//设计用来pipeline方式调用
func (node *TreeNode) SetVal(val int) *TreeNode {
	node.Value = val
	//因为返回来对应的实例来实现链式调用：node.SetVal(5).GetVal()//这样就实现包内方法的链式调用
	return node
}
