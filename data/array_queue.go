package data

import "fmt"

type ArrayQueue struct {
	items         []interface{}
	n, head, tail int
}

func NewArrayQueue(capacity int) *ArrayQueue {
	return &ArrayQueue{
		items: make([]interface{}, capacity),
		n:     capacity,
	}
}

func (q *ArrayQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

func (q *ArrayQueue) IsFull() bool {
	if q.tail == q.n {
		return true
	}
	return false
}

func (q *ArrayQueue) EnQueue(v interface{}) bool {
	if q.IsFull() {
		if q.head == 0 {
			// tail==n && head==0 表示队列满了
			return false
		}
		// 数据搬移
		for i := q.head; i < q.tail; i++ {
			q.items[i-q.head] = q.items[i]
		}

		q.tail -= q.head
		q.head = 0
	}
	q.items[q.tail] = v
	q.tail++
	return true
}

func (q *ArrayQueue) DeQueue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	v := q.items[q.head]
	q.head++
	return v, true
}

func (q *ArrayQueue) String() string {
	if q.IsEmpty() {
		return "empty queue"
	}
	result := "head"
	for i := q.head; i <= q.tail-1; i++ {
		result += fmt.Sprintf("<-%+v", q.items[i])
	}
	result += "<-tail"
	return result
}
