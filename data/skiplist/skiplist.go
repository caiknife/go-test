package skiplist

import (
	"math"
	"fmt"
	"math/rand"
)

const MAX_LEVEL = 16 //最高层数

type List interface {
	Length() int
	Level() int
	String() string
}

// 跳表节点结构体
type SkipListNode struct {
	v        interface{}     // 跳表保存的值
	score    int             // 用于排序的分值
	level    int             // 层高
	forwards []*SkipListNode // 每层前进指针
}

func (this *SkipListNode) String() string {
	return fmt.Sprintf("value:%+v, score:%+v, level:%+v, forwards:%+v", this.v, this.score, this.level, this.forwards)
}

// 跳表结构体
type SkipList struct {
	List
	head   *SkipListNode // 跳表头结点
	level  int           // 跳表当前层数
	length int           // 跳表长度
}

// 实例化化跳表节点对象
func NewSkipListNode(v interface{}, score int, level int) *SkipListNode {
	return &SkipListNode{
		v:        v,
		score:    score,
		level:    level,
		forwards: make([]*SkipListNode, level, level),
	}
}

// 实例化跳表对象
func NewSkipList() *SkipList {
	head := NewSkipListNode(0, math.MaxUint32, MAX_LEVEL)
	return &SkipList{head: head, level: 1, length: 0,}
}

// 跳表长度
func (this *SkipList) Length() int {
	return this.length
}

// 跳表层级
func (this *SkipList) Level() int {
	return this.level
}

// 跳表输出
func (this *SkipList) String() string {
	return fmt.Sprintf("level:%+v, length:%+v", this.level, this.length)
}

// 插入节点到跳表中
func (this *SkipList) Insert(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找插入位置
	cur := this.head
	update := [MAX_LEVEL]*SkipListNode{}
	for i := MAX_LEVEL - 1; i >= 0; i-- {
		for cur.forwards[i] != nil {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
		if cur.forwards[i] == nil {
			update[i] = cur
		}
	}

	// 通过随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}

	// 创建一个新的跳表节点
	newNode := NewSkipListNode(v, score, level)

	// 原有节点连接
	for i := 0; i <= level-1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > this.level {
		this.level = level
	}

	//更新跳表长度
	this.length++

	return 0
}

// 查找
func (this *SkipList) Find(v interface{}, score int) *SkipListNode {
	if v == nil || this.length == 0 {
		return nil
	}

	cur := this.head
	for i := this.level - 1; i >= 0; i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				return cur.forwards[i]
			} else if cur.forwards[i].score > score {
				break
			}
			cur = cur.forwards[i]
		}
	}

	return nil
}

// 删除节点
func (this *SkipList) Delete(v interface{}, score int) int {
	if v == nil {
		return 1
	}

	// 查找前驱节点
	cur := this.head
	// 记录前驱路径
	update := [MAX_LEVEL]*SkipListNode{}
	for i := this.level - 1; i >= 0; i-- {
		update[i] = this.head
		for nil != cur.forwards[i] {
			if cur.forwards[i].score == score && cur.forwards[i].v == v {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		}
	}

	cur = update[0].forwards[0]
	for i := cur.level - 1; i >= 0; i-- {
		if update[i] == this.head && cur.forwards[i] == nil {
			this.level = i
		}

		if nil == update[i].forwards[i] {
			update[i].forwards[i] = nil
		} else {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
	}

	this.length--

	return 0
}
