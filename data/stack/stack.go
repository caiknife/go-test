package stack

type Stack interface {
	Push(v interface{})
	Pop()
	IsEmpty() bool
	Top() interface{}
	Flush()
	String() string
	Print()
}
