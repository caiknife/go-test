package queue

type Queue interface {
	DeQueue() (interface{}, bool)
	EnQueue(v interface{}) bool
	IsEmpty() bool
	String() string
}
