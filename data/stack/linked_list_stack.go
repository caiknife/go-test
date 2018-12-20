package stack

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedListStack struct {
	top *Node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (s *LinkedListStack) IsEmpty() bool {
	if s.top == nil {
		return true
	}
	return false
}

func (s *LinkedListStack) Push(v interface{}) {
	s.top = &Node{
		next:  s.top,
		value: v,
	}
}

func (s *LinkedListStack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}

	v := s.top.value
	s.top = s.top.next
	return v
}

func (s *LinkedListStack) Top() interface{} {
	if s.IsEmpty() {
		return nil
	}
	return s.top.value
}

func (s *LinkedListStack) Flush() {
	s.top = nil
}

func (s *LinkedListStack) String() string {
	if s.IsEmpty() {
		return "empty queue"
	}

	result := "top"
	for cur := s.top; cur != nil; cur = cur.next {
		result += fmt.Sprintf("->%+v", cur.value)
	}
	result += "->bottom"
	return result
}

func (s *LinkedListStack) Print() {
	fmt.Println(s.String())
}
