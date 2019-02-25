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
		if compareResult > 0 {
			//v > nodeV
			p = p.right
		} else if compareResult < 0 {
			//v < nodeV
			p = p.left
		} else {
			return p
		}
	}
	return nil
}

func (this *BinarySearchTree) Insert(v interface{}) bool {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			if p.right == nil {
				p.right = NewNode(v)
				break
			}
			p = p.right
		} else if compareResult < 0 {
			if p.left == nil {
				p.left = NewNode(v)
				break
			}
			p = p.left
		} else {
			return false
		}
	}
	return true
}

func (this *BinarySearchTree) Delete(v interface{}) bool {
	var pp *TreeNode = nil
	p := this.root
	deleteLeft := false

	for p != nil {
		compareResult := this.compareFunc(v, p.data)
		if compareResult > 0 {
			pp = p
			p = p.right
			deleteLeft = false
		} else if compareResult < 0 {
			pp = p
			p = p.left
			deleteLeft = true
		} else {
			break
		}
	}

	if p == nil {
		// 需要删除的节点不存在
		return false
	} else if p.left == nil && p.right == nil {
		// 删除的是一个叶子节点
		if pp == nil {
			if deleteLeft {
				pp.left = nil
			} else {
				pp.right = nil
			}
		} else {
			// 根节点
			this.root = nil
		}
	} else if p.right != nil {
		// 删除的是一个有右子节点，不一定有左子节点的节点
		// 找到p节点右子节点的最小节点
		pq := p
		q := p.right // 向右走一步
		fromRight := true
		for q.left != nil { //向左走到底
			pq = q
			q = q.left
			fromRight = false
		}
		if fromRight {
			pq.right = nil
		} else {
			pq.left = nil
		}
		q.left = p.left
		q.right = p.right
		if pp == nil { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				pq.left = q
			} else {
				pq.right = q
			}
		}
	} else { //删除的是一个只有左孩子的节点
		if nil != pp {
			if deleteLeft {
				pp.left = p.left
			} else {
				pp.right = p.left
			}
		} else {
			if deleteLeft {
				this.root = p.left
			} else {
				this.root = p.left
			}
		}
	}

	return true
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
