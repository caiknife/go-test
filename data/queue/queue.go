package queue

type Queue interface {
	DeQueue() (interface{}, bool)
	EnQueue() bool
	IsEmpty() bool
	String() string
}
