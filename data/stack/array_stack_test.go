package stack

import "testing"

func TestArrayStackPush(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	t.Log(s.Pop())
	s.Push(3)
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Push(4)
	t.Log(s.Pop())
	s.Print()
}

func TestArrayStackPop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Print()

	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	t.Log(s.Pop())
	s.Print()
}

func TestArrayStackTop(t *testing.T) {
	s := NewArrayStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
	t.Log(s.Top())
	s.Pop()
}
