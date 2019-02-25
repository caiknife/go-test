package tree

import "testing"

var compareFunc = func(v, nodeV interface{}) int {
	vInt := v.(int)
	nodeVInt := nodeV.(int)

	if vInt > nodeVInt {
		return 1
	} else if vInt < nodeVInt {
		return -1
	}

	return 0
}

func TestBinarySearchTree_Find(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Find(2))
}

func TestBinarySearchTree_Insert(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	bst.PreOrder()
	bst.InOrder()
	bst.PostOrder()
}

func TestBinarySearchTree_Max(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Max())
}

func TestBinarySearchTree_Min(t *testing.T) {
	bst := NewBST(1, compareFunc)

	bst.Insert(3)
	bst.Insert(1)
	bst.Insert(2)
	bst.Insert(7)
	bst.Insert(5)

	t.Log(bst.Min())
}
