package tree

import "fmt"

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

func NewNode(data interface{}) *TreeNode {
	return &TreeNode{data: data}
}

func (this *TreeNode) String() string {
	return fmt.Sprintf("v:%+v, left:%+v, right:%+v", this.data, this.left, this.right)
}
