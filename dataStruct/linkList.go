package dataStruct

import "errors"

/*
 * @Description 实现一个双端、无环链表
 * @Author: zyz
 * @Date: 2024/4/19 下午 11:24
 */

type listNode[T any] struct {
	//前置节点
	prev *listNode[T]
	//后置节点
	next *listNode[T]
	//节点的值
	value T
}

type List[T any] struct {
	//链表的头节点
	head *listNode[T]
	//链表的尾节点
	tail *listNode[T]
	//链表的长度
	len int
	//节点复制函数
	copy func(T) T
	//节点值对比函数
	equal func(T, T) bool
}

// 创建一个新的链表
func newList[T any]() *List[T] {
	return &List[T]{}
}

// 创建一个新的链表节点
func newListNode[T any](value T) *listNode[T] {
	return &listNode[T]{value: value}
}

// 创建一个新的链表空节点
func newListEmptyNode[T any]() *listNode[T] {
	return &listNode[T]{}
}

// 在链表的头部插入一个节点
func (l *List[T]) PushFront(value T) {
	node := newListEmptyNode[T]()
	node.value = value
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
}

// 在链表的尾部插入一个节点
func (l *List[T]) PushBack(value T) {
	node := newListEmptyNode[T]()
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

// 在指定节点前后插入一个新节点
func (l *List[T]) InsertBefore(node *listNode[T], newNode *listNode[T]) {
	newNode.prev = node.prev
	newNode.next = node
	if node.prev == nil {
		l.head = newNode
	} else {
		node.prev.next = newNode
	}
	node.prev = newNode
	l.len++
}

// 删除指定的节点
func (l *List[T]) Remove(node *listNode[T]) {
	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
	l.len--
}

// 复制整个链表
func (l *List[T]) Copy() (*List[T], error) {
	if l.copy == nil {
		return nil, errors.New("copy function is nil")
	}
	newList := newList[T]()
	newList.copy = l.copy
	newList.equal = l.equal
	for i := l.head; i != nil; i = i.next {
		newList.PushBack(l.copy(i.value))
	}
	return newList, nil
}

// 在链表中搜索指定的键
func (l *List[T]) Search(key T) (*listNode[T], error) {

	if l.equal == nil {
		return nil, errors.New("equal function is nil")
	}

	for i := l.head; i != nil; i = i.next {
		if l.equal(i.value, key) {
			return i, nil
		}
	}
	return nil, nil
}

// 返回链表中指定索引的节点
func (l *List[T]) Index(index int) *listNode[T] {
	if index < 0 || index >= l.len {
		return nil
	}
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}

// 将链表的尾部节点移动到头部
func (l *List[T]) Rotate() {
	if l.tail == nil {
		return
	}
	l.tail.next = l.head
	l.head.prev = l.tail
	l.head = l.tail
	l.tail = l.tail.prev
	l.head.prev = nil
	l.tail.next = nil

}

// 将链表的头部节点移动到尾部
func (l *List[T]) RotateBack() {
	if l.head == nil {
		return
	}
	l.head.prev = l.tail
	l.tail.next = l.head
	l.tail = l.head
	l.head = l.head.next
	l.head.prev = nil
	l.tail.next = nil
}

// 将两个链表连接在一起
func (l *List[T]) Concat(list *List[T]) {
	if list == nil || list.head == nil {
		return
	}

	if l.tail == nil {
		l.head = list.head
		l.tail = list.tail
	} else {
		l.tail.next = list.head
		list.head.prev = l.tail
		l.tail = list.tail
	}
	l.len += list.len
}

// 将一个节点链接到链表的头部
func (l *List[T]) Append(node *listNode[T]) {
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.len++
}

// 将一个节点链接到链表的尾部
func (l *List[T]) AppendBack(node *listNode[T]) {
	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		node.prev = l.tail
		l.tail.next = node
		l.tail = node
	}
	l.len++
}

// 从链表中断开一个节点
func (l *List[T]) Detach(node *listNode[T]) {
	if node.prev == nil {
		l.head = node.next
	} else {
		node.prev.next = node.next
	}
	if node.next == nil {
		l.tail = node.prev
	} else {
		node.next.prev = node.prev
	}
	l.len--
}
