package main

import "fmt"

type treeNode struct {
	val   int
	left  *treeNode
	right *treeNode
}

func creatTreeNode(val int) *treeNode {
	return &treeNode{val: val} //返回的引用是当前函数局部变量【显然也是可以，闭包既然可以实现】
}

func (node treeNode) print() {
	//
}

func (node *treeNode) setNodeVal(val int) {
	//修改值使用指针接收者
}

func main() {

	var root treeNode

	root.val = 3

	root.left = new(treeNode)

	root.right = &treeNode{6, nil, nil}

	fmt.Println(root)
	fmt.Println(root.right)

	root.right.right = creatTreeNode(15)
	fmt.Println(root.right.right)
}
