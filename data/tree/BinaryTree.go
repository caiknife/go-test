package tree

import (
	"go-test/data/stack"
	"fmt"
)

type BinaryTree struct {
	root *TreeNode
}

func NewBinaryTree(v interface{}) *BinaryTree {
	return &BinaryTree{NewNode(v)}
}

func (this *BinaryTree) PreOrder() {
	p := this.root
	s := stack.NewArrayStack()
}

func (this *BinaryTree) InOrder() {
	p := this.root
	s := stack.NewArrayStack()
}

func (this *BinaryTree) PostOrder() {
	s1 := stack.NewArrayStack()
	s2 := stack.NewArrayStack()
}

func PreOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	fmt.Sprintf("%+v ", p.data)
	PreOrderTraverse(p.left)
	PreOrderTraverse(p.right)
}

func InOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	InOrderTraverse(p.left)
	fmt.Sprintf("%+v ", p.data)
	InOrderTraverse(p.right)
}

func PostOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	PostOrderTraverse(p.left)
	PostOrderTraverse(p.right)
	fmt.Sprintf("%+v ", p.data)
}
