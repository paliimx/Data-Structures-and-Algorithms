package StackLinkedList

import "cmp"

type Node[T cmp.Ordered] struct {
	data T
	next *Node[T]
}

type Stack[T cmp.Ordered] struct {
	top *Node[T]
}

func (list *Stack[T]) Push(i T) {
	data := &Node[T]{data: i}
	if list.top != nil {
		data.next = list.top
	}
	list.top = data
}

func (list *Stack[T]) Pop() (T, bool) {
	var t T
	if list.top == nil {
		return t, false
	}
	i := list.top.data
	list.top = list.top.next
	return i, true
}

func (list *Stack[T]) Peek() (T, bool) {
	var t T
	if list.top == nil {
		return t, false
	}
	return list.top.data, true
}

func (list *Stack[T]) Get() []T {

	var items []T

	current := list.top
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Stack[T]) IsEmpty() bool {
	return list.top == nil
}

func (list *Stack[T]) Empty() {
	list.top = nil
}
