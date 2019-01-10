package linkedlist

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedList struct {
	head   *ListNode
	length uint
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func (list *ListNode) GetNext() *ListNode {
	return list.next
}

func (list *ListNode) GetValue() interface{} {
	return list.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (list *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	list.length++
	return true
}

//在某个节点前面插入节点
func (list *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == list.head {
		return false
	}
	cur := list.head.next
	pre := list.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	list.length++
	return true
}

//在链表头部插入节点
func (list *LinkedList) InsertToHead(v interface{}) bool {
	return list.InsertAfter(list.head, v)
}

//在链表尾部插入节点
func (list *LinkedList) InsertToTail(v interface{}) bool {
	cur := list.head
	for nil != cur.next {
		cur = cur.next
	}
	return list.InsertAfter(cur, v)
}

//通过索引查找节点
func (list *LinkedList) FindByIndex(index uint) *ListNode {
	if index >= list.length {
		return nil
	}
	cur := list.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

//删除传入的节点
func (list *LinkedList) DeleteNode(p *ListNode) bool {
	if nil == p {
		return false
	}
	cur := list.head.next
	pre := list.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	list.length--
	return true
}

//打印链表
func (list *LinkedList) Print() {
	fmt.Println(list.String())
}

func (list *LinkedList) String() string {
	cur := list.head.next
	format := ""
	for cur != nil {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil {
			format += "->"
		}
	}
	return format
}

// 链表反转 复杂度O(n)
func (list *LinkedList) Reverse() {
	// 不需要反转
	if list.head == nil || list.head.next == nil || list.head.next.next == nil {
		return
	}

	var pre *ListNode = nil
	cur := list.head.next
	// 遍历链表
	for cur != nil {
		var tmp = cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	list.head.next = pre
}

// 判断链表是否有环
func (list *LinkedList) HasCycle() bool {
	if list.head != nil {
		slow := list.head
		fast := list.head

		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next
			if slow == fast {
				return true
			}
		}
	}
	return false
}
