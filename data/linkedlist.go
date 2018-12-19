package data

import "fmt"

// 链表节点
type ListNode struct {
	next  *ListNode
	value interface{}
}

// 链表开始
type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

func (l *ListNode) GetNext() *ListNode {
	return l.next
}

func (l *ListNode) GetValue() interface{} {
	return l.value
}

// 在某个节点之后插入新节点
func (l *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if p == nil {
		return false
	}

	cur := l.head.next

	for cur != nil {
		if cur == p {
			break
		}
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	oldNext := cur.next
	cur.next = newNode
	newNode.next = oldNext
	l.length++
	return true
}

// 在某个节点之前插入新节点
func (l *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if p == nil || p == l.head {
		return false
	}

	cur := l.head.next
	pre := l.head

	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	l.length++
	return true
}

// 在链表头部插入新节点
func (l *LinkedList) InsertAtHead(v interface{}) bool {
	return l.InsertAfter(l.head, v)
}

// 在链表尾部插入新节点
func (l *LinkedList) InsertAtTail(v interface{}) bool {
	cur := l.head
	for cur.next != nil {
		cur = cur.next
	}
	return l.InsertAfter(cur, v)
}

// 通过索引查找节点
func (l *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= l.length {
		return nil
	}

	cur := l.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除节点
func (l *LinkedList) DeleteNode(p *ListNode) bool {
	if p == nil {
		return false
	}

	cur := l.head.next
	pre := l.head

	for cur != nil {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}

	if cur == nil {
		return false
	}

	pre.next = p.next
	p = nil
	l.length--
	return true
}

// 打印链表
func (l *LinkedList) Print() {
	cur := l.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	fmt.Println(format)
}
