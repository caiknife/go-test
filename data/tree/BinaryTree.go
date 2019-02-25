package tree

import (
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
	PreOrderTraverse(p)
	fmt.Println("")
}

func (this *BinaryTree) InOrder() {
	p := this.root
	InOrderTraverse(p)
	fmt.Println("")
}

func (this *BinaryTree) PostOrder() {
	p := this.root
	PostOrderTraverse(p)
	fmt.Println("")
}

func PreOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	fmt.Printf("%+v ", p.data)
	PreOrderTraverse(p.left)
	PreOrderTraverse(p.right)
}

func InOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	InOrderTraverse(p.left)
	fmt.Printf("%+v ", p.data)
	InOrderTraverse(p.right)
}

func PostOrderTraverse(node *TreeNode) {
	p := node
	if p == nil {
		return
	}
	PostOrderTraverse(p.left)
	PostOrderTraverse(p.right)
	fmt.Printf("%+v ", p.data)
}
