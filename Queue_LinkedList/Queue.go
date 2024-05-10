package QueueLinkedList

import "cmp"

type Node[T cmp.Ordered] struct {
	data T
	next *Node[T]
}

type Queue[T cmp.Ordered] struct {
	rear *Node[T]
}

func (list *Queue[T]) Enqueue(i T) {
	data := &Node[T]{data: i}
	if list.rear != nil {
		data.next = list.rear
	}
	list.rear = data
}

func (list *Queue[T]) Dequeue() (T, bool) {
	var t T
	if list.rear == nil {
		return t, false
	}
	if list.rear.next == nil {
		i := list.rear.data
		list.rear = nil
		return i, true
	}
	current := list.rear
	for {
		if current.next.next == nil {
			i := current.next.data
			current.next = nil
			return i, true
		}
		current = current.next
	}
}

func (list *Queue[T]) Peek() (T, bool) {
	var t T
	if list.rear == nil {
		return t, false
	}
	return list.rear.data, true
}

func (list *Queue[T]) Get() []T {
	var items []T
	current := list.rear
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *Queue[T]) IsEmpty() bool {
	return list.rear == nil
}

func (list *Queue[T]) Empty() {
	list.rear = nil
}
