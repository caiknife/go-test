package tree

type BinarySearchTree struct {
	*BinaryTree
	compareFunc func(v, nodeV interface{}) int
}

func NewBST(rootV interface{}, compareFunc func(v, nodeV interface{}) int) *BinarySearchTree {
	if compareFunc == nil {
		return nil
	}
	return &BinarySearchTree{BinaryTree: NewBinaryTree(rootV), compareFunc: compareFunc}
}

func (this *BinarySearchTree) Find(v interface{}) *TreeNode {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return p
		} else if compareResult < 0 { // v < p.data
			p = p.left
		} else { // v > p.data
			p = p.right
		}
	}
	return nil
}

func (this *BinarySearchTree) Insert(v interface{}) bool {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult == 0 {
			return false
		} else if compareResult < 0 {
			if p.left == nil {
				p.left = NewNode(v)
				break
			}
			p = p.left
		} else {
			if p.right == nil {
				p.right = NewNode(v)
				break
			}
		}
	}

	return true
}

func (this *BinarySearchTree) Delete(v interface{}) bool {
	// TODO
}

func (this *BinarySearchTree) Max() *TreeNode {
	p := this.root
	for p.right != nil {
		p = p.right
	}
	return p
}

func (this *BinarySearchTree) Min() *TreeNode {
	p := this.root
	for p.left != nil {
		p = p.left
	}
	return p
}
