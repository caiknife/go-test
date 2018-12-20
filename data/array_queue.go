package data

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
		return false
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
