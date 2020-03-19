package linkedlist

// 双向链表
type LinkedList struct {
	linkedNode
}

// 双向链表构造函数
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// 添加元素
func (l2 *LinkedList) Add(index int, args ...interface{}) error {
	// 1. 第一次添加
	if l2.size == 0 {
		l2.first = &nodes{}
		l2.first.data = args[0]
		l2.last = &nodes{}
		l2.last.data = l2.first.data
		l2.size++
		if len(args) > 1 {
			err := l2.append(args[1:], l2.size-1)
			if err != nil {
				l2.Clear()
				return err
			}
		}
		return nil
	}
	// 2. 尾部追加
	if l2.size == index {
		err := l2.append(args, l2.size-1)
		return err
	}
	// 3. 中间插入
	if tmpIndex := l2.size - 1; index > 0 && index <= tmpIndex {
		// 3.1 获取要插入的节点，先保存
		nextNode, err := l2.node(index)
		if err != nil {
			return err
		}
		// 3.1.1 获取插入位置的上一个节点
		frontNode := nextNode.prev
		for i, v := range args {
			// 1.1 构建节点
			newNode := &nodes{}
			newNode.data = v
			// 1.2 构建节点指向步骤1.的节点
			newNode.prev = frontNode
			if i == len(args)-1 {
				newNode.next = nextNode
				nextNode.prev = newNode
			}
			// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
			frontNode.next = newNode
			// 1.4 把当前的保存，让下一次循环变成父节点
			frontNode = newNode
			l2.size++
		}
		return nil
	}
	// 4. 第一位插入
	// 4.1 获取要首节点，先保存
	nextNode := l2.first
	frontNode := &nodes{}
	for i, v := range args {
		// 1.1 构建节点
		newNode := &nodes{}
		newNode.data = v
		newNode.prev = frontNode
		// 1.2 构建的是第一个，prev 指向最后一个节点，
		// 最后一个节点的 enxt 指向入口节点第一个,形成双闭环
		if i == 0 {
			newNode.prev = nil
			l2.first = newNode
		} else {
			if i == len(args)-1 {
				newNode.next = nextNode
				nextNode.prev = newNode
			}
		}
		// 1.3 步骤1.的节点的下一个节点指向刚构建的节点
		frontNode.next = newNode
		// 1.4 把当前的保存，让下一次循环变成父节点
		frontNode = newNode
		l2.size++
	}
	return nil
}

func (l2 *LinkedList) append(args []interface{}, index int, a ...interface{}) error {
	nextNode, err := l2.node(index) // 获取节点
	if err != nil {
		return err
	}
	isLastNode := len(a) < 1
	endNode := &nodes{}
	for i, v := range args {
		newNode := &nodes{}
		newNode.data = v
		newNode.prev = nextNode
		nextNode.next = newNode // 把当前的节点赋值父节点
		nextNode = newNode      // 把当前的保存，让下一次循环变成父节点
		l2.size++
		if i == len(args)-1 && isLastNode {
			endNode = newNode
		}
	}
	if isLastNode {
		l2.last = endNode
	}
	return nil
}

// 清除链表
func (l2 *LinkedList) Clear() {
	l2.first = nil
	l2.size = 0
	return
}

// 获取元素
func (l2 *LinkedList) Get(index int) (interface{}, error) {
	data, err := l2.node(index)
	if err != nil {
		return nil, err
	}
	return data.data, nil
}

// 获取元素对应的第一次的索引，暂时不支持获取引用类型元素的获取
func (l2 *LinkedList) IndexOf(element interface{}) int {
	return indexOf(element, &l2.linkedNode)
}

// 获取元素的容量
func (l2 *LinkedList) Size() int {
	return l2.size
}

// 根据索引设置元素
func (l2 *LinkedList) Set(index int, element interface{}) (interface{}, error) {
	node, err := l2.node(index)
	if err != nil {
		return nil, err
	}
	node.data = element
	return element, err
}

// 移除元素
func (l2 *LinkedList) Remove(index int) (interface{}, error) {
	err := rangeCheck(index, l2.size)
	if err != nil {
		return nil, err
	}
	// 0. 容量为1，直接清空
	if l2.size == 1 {
		element := l2.first.data
		l2.Clear()
		return element, nil
	}
	// 1. 移除第一个节点
	if index == 0 {
		// 0. 保存被移除的元素
		element := l2.first.data
		// 1.1 获取第二个节点，先保存
		// 1.2 入口元素指向 1.1 的元素
		node := l2.first.next
		node.prev = nil
		l2.first = node
		l2.size--
		return element, nil
	}
	// 2. 移除最后的元素
	if index == l2.size-1 {
		// 0. 保存被移除的元素
		element := l2.last.data
		node := l2.last.prev
		node.next = nil
		l2.last = node
		l2.size--
		return element, nil
	}
	// 3. 移除中间的元素
	// 3.1. 保存被移除的元素
	node, err := l2.node(index)
	if err != nil {
		return nil, err
	}
	// 3.2 移除的节点前一个
	firstNode := node.prev
	// 3.3 移除的节点后一个
	tailNode := node.next
	firstNode.next = tailNode
	tailNode.prev = firstNode
	l2.size--
	return node.data, nil
}

// 获取index位置对应的节点对象
func (l2 *LinkedList) node(index int) (*nodes, error) {
	err := rangeCheck(index, l2.size)
	if err != nil {
		return nil, err
	}
	// 1. > size 的一半，尾部查找
	if index <= (l2.size >> 1) {
		node := l2.first
		for i := 0; i < index; i++ {
			node = node.next
		}
		return node, nil
	} else {
		node := l2.last
		for i := l2.size - 1; i > index; i-- {
			node = node.prev
		}
		return node, nil
	}
}
