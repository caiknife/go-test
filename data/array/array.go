package array

import (
	"errors"
	"fmt"
)

type Object interface {
	Len() int
	IsIndexOutOfRange(index int) bool
	Find(index int) (interface{}, error)
	Insert(index int, v interface{}) error
	InsertToTail(v interface{}) error
	Delete(index int) (interface{}, error)
	String() string
	Print()
}

type Array struct {
	data   []interface{}
	length int
	Object
}

func NewArray(capacity int) *Array {
	if capacity == 0 {
		return nil
	}

	return &Array{
		data:   make([]interface{}, capacity, capacity),
		length: 0,
	}
}

func (a *Array) Len() int {
	return a.length
}

func (a *Array) IsIndexOutOfRange(index int) bool {
	if index >= cap(a.data) {
		return true
	}
	return false
}

func (a *Array) Find(index int) (interface{}, error) {
	if a.IsIndexOutOfRange(index) {
		return nil, errors.New("index out of range")
	}
	return a.data[index], nil
}

// 插入元素到数组指定元素
func (a *Array) Insert(index int, v interface{}) error {
	if a.Len() == cap(a.data) {
		return errors.New("full array")
	}
	if index != a.length && a.IsIndexOutOfRange(index) {
		return errors.New("index out of range")
	}

	for i := a.length; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = v
	a.length++
	return nil
}

func (a *Array) InsertToTail(v interface{}) error {
	return a.Insert(a.Len(), v)
}

func (a *Array) Delete(index int) (interface{}, error) {
	if a.IsIndexOutOfRange(index) {
		return nil, errors.New("index out of range")
	}
	v := a.data[index]
	for i := index; i < a.Len()-1; i++ {
		a.data[i] = a.data[i+1]
	}
	a.length--
	return v, nil
}

func (a *Array) String() string {
	var result string
	for i := 0; i < a.Len(); i++ {
		result += fmt.Sprintf("|%+v", a.data[i])
	}
	return result
}

func (a *Array) Print() {
	fmt.Println(a.String())
}
