package skiplist

import (
	"math"
	"fmt"
)

const MAX_LEVEL = 16 //最高层数

type List interface {
	Length() int
	Level() int
	String() string
}

type SkipListNode struct {
	v        interface{}     // 跳表保存的值
	score    int             // 用于排序的分值
	level    int             // 层高
	forwards []*SkipListNode // 每层前进指针
}

type SkipList struct {
	List
	head   *SkipListNode // 跳表头结点
	level  int           // 跳表当前层数
	length int           // 跳表长度
}

func NewSkipListNode(v interface{}, score int, level int) *SkipListNode {
	return &SkipListNode{
		v:        v,
		score:    score,
		level:    level,
		forwards: make([]*SkipListNode, level, level),
	}
}

func NewSkipList() *SkipList {
	head := NewSkipListNode(0, math.MaxUint32, MAX_LEVEL)
	return &SkipList{head: head, level: 1, length: 0,}
}

func (this *SkipList) Length() int {
	return this.length
}

func (this *SkipList) Level() int {
	return this.level
}

func (this *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", this.level, this.length)
}
