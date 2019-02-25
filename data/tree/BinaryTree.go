package tree

import "fmt"

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{NewNode(v)}
}
