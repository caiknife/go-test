package data

type CircularQueue struct {
	items         []interface{}
	n, head, tail int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		items: make([]interface{}, capacity),
		n:     capacity,
	}
}

func (q *CircularQueue) IsEmpty() bool {
	if q.head == q.tail {
		return true
	}
	return false
}

func (q *CircularQueue) IsFull() bool {
	if (q.tail+1)%q.n == q.head {
		return true
	}
	return false
}

func (q *CircularQueue) EnQueue(item interface{}) bool {
	if q.IsFull() {
		return false
	}

	q.items[q.tail] = item
	q.tail = (q.tail + 1) % q.n
	return true
}

func (q *CircularQueue) DeQueue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	item := q.items[q.head]
	q.head = (q.head + 1) % q.n
	return item, true
}
