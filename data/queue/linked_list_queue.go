package queue

import "fmt"

type ListNode struct {
	next  *ListNode
	value interface{}
}

type LinkedListQueue struct {
	head   *ListNode
	tail   *ListNode
	length int
}

func NewLinkedListQueue() *LinkedListQueue {
	return &LinkedListQueue{nil, nil, 0}
}

func (q *LinkedListQueue) EnQueue(v interface{}) bool {
	n := &ListNode{nil, v}

	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		q.tail.next = n
		q.tail = n
	}
	return true
}

func (q *LinkedListQueue) IsEmpty() bool {
	return q.head == nil
}

func (q *LinkedListQueue) DeQueue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	v := q.head.value
	q.head = q.head.next
	q.length--
	return v, true
}

func (q *LinkedListQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	for cur := q.head; cur != nil; cur = cur.next {
		result += fmt.Sprintf("<-%+v", cur.value)
	}
	result += "<-tail"
	return result
}
