package linkedlist

import (
	"fmt"
	"reflect"
)

// 存储的元素
type nodeData struct {
	data interface{}
}

// 单向链表节点
type node struct {
	next *node // 下一个节点指针
	nodeData
}

// 单向链表结构体
type singleLinkedNode struct {
	size  int   // 元素的长度
	first *node // 入口节点
}

// 双向链表节点
type nodes struct {
	prev *nodes
	next *nodes
	nodeData
}

// 双向链表结构体
type linkedNode struct {
	size  int    // 元素的长度
	first *nodes // 头入口节点
	last  *nodes //尾入口节点
}

const (
	ELEMENT_NOT_FOUND int = -1
)

// 范围检查
func rangeCheck(index, size int) (err error) {
	if index < 0 || index >= size {
		err = fmt.Errorf("Index cross boundary Index:%d , Cap:%d\n", index, size)
	}
	return
}

// ---------------单向链表共用方法------------------

func singleIndexOf(element interface{}, ls *singleLinkedNode) int {
	if element == nil {
		node := ls.first
		for i := 0; i < ls.size; i++ {
			if node.data == nil {
				return i
			}
		}
		node = node.next
	} else {
		node := ls.first
		for i := 0; i < ls.size; i++ {
			if reflect.DeepEqual(element, node.data) {
				return i
			}
			node = node.next
		}
		node = node.next
	}
	return ELEMENT_NOT_FOUND
}

// ---------------双向链表共用方法------------------
func indexOf(element interface{}, ls *linkedNode) int {
	node1 := ls.first
	node2 := ls.last
	lastNum := ls.size
	svg := (ls.size - 1) >> 1 // 边界
	if element == nil {
		for i := 0; i < ls.size; i++ {
			if i > svg {
				break
			}
			lastNum--
			if node1.data == nil {
				return i
			}
			if node2.data == nil {
				return lastNum
			}
			node1 = node1.next
			node2 = node2.prev
		}
	} else {
		for i := 0; i < ls.size; i++ {
			if i > svg {
				break
			}
			lastNum--
			if reflect.DeepEqual(element, node1.data) {
				return i
			}
			if reflect.DeepEqual(element, node2.data) {
				return lastNum
			}
			node1 = node1.next
			node2 = node2.prev
		}
	}
	return ELEMENT_NOT_FOUND
}
