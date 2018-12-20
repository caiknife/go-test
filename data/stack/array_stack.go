package stack

import "fmt"

type Stack interface {
	Push(v interface{})
	Pop()
	IsEmpty() bool
	Top() interface{}
	Flush()
	String() string
	Print()
}

type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	var array = make([]interface{}, 0, 32)
	return &ArrayStack{
		data: array,
		top:  -1,
	}
}

func (s *ArrayStack) IsEmpty() bool {
	if s.top < 0 {
		return true
	}
	return false
}

func (s *ArrayStack) Push(v interface{}) {
	s.top += 1

	if s.top > len(s.data)-1 {
		s.data = append(s.data, v)
	} else {
		s.data[s.top] = v
	}
}

func (s *ArrayStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	v := s.data[s.top]
	s.top--
	return v
}

func (s *ArrayStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.data[s.top]
}

func (s *ArrayStack) Flush() {
	s.top = -1
}

func (s *ArrayStack) String() string {
	if s.IsEmpty() {
		return "empty stack"
	}
	result := "top"
	for i := s.top; i > -1; i-- {
		result += fmt.Sprintf("->%+v", s.data[i])
	}
	result += "bottom"
	return result
}

func (s *ArrayStack) Print() {
	fmt.Println(s.String())
}
