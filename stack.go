package stack

import "stack/linkedlist"

type Stack struct {
	*linkedlist.LinkedList
}

// 初始化
func NewInit() *Stack {
	tmp := new(Stack)
	tmp.LinkedList = linkedlist.NewLinkedList()
	return tmp
}

// 清空栈
func (ls Stack) Clear() {
	ls.LinkedList.Clear()
	return
}

// 是否为空
func (ls Stack) IsEmpty() bool {
	return ls.LinkedList.Size() == 0
}

// 元素的数量
func (ls Stack) Size() int {
	return ls.LinkedList.Size()
}

// 入栈
func (ls Stack) Push(e ...interface{}) error {
	return ls.LinkedList.Add(ls.LinkedList.Size(), e...)
}

// 出栈
func (ls Stack) Pop() (interface{}, error) {
	return ls.LinkedList.Remove(ls.LinkedList.Size() - 1)
}

// 获取栈顶元素
func (ls Stack) Top() (interface{}, error) {
	return ls.LinkedList.Get(ls.LinkedList.Size() - 1)
}
