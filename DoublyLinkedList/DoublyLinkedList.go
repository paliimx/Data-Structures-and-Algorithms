package DoublyLinkedList

import "cmp"

type Node[T cmp.Ordered] struct {
	data T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T cmp.Ordered] struct {
	head *Node[T]
	tail *Node[T]
}

func (list *LinkedList[T]) InsertFirst(i T) {
	data := &Node[T]{data: i}
	if list.head != nil {
		list.head.prev = data
		data.next = list.head
	}
	list.head = data
}

func (list *LinkedList[T]) InsertLast(i T) {
	data := &Node[T]{data: i}
	if list.head == nil {
		list.head = data
		list.tail = data
		return
	}
	if list.tail != nil {
		list.tail.next = data
		data.prev = list.tail
	}
	list.tail = data
}

func (list *LinkedList[T]) RemoveByValue(i T) bool {
	if list.head == nil {
		return false
	}
	if cmp.Compare(list.head.data, i) == 0 {
		list.head = list.head.next
		list.head.prev = nil
		return true
	}
	if cmp.Compare(list.tail.data, i) == 0 {
		list.tail = list.tail.prev
		list.tail.next = nil
		return true
	}
	current := list.head
	for current.next != nil {
		if cmp.Compare(current.next.data, i) == 0 {
			if current.next.next != nil {
				current.next.next.prev = current
			}
			current.next = current.next.next
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList[T]) RemoveByIndex(i int) bool {
	if list.head == nil {
		return false
	}
	if i < 0 {
		return false
	}
	if i == 0 {
		list.head.prev = nil
		list.head = list.head.next
		return true
	}
	current := list.head
	for u := 1; u < i; u++ {
		if current.next.next == nil {
			return false
		}
		current = current.next
	}
	if current.next.next != nil {
		current.next.next.prev = current
	}
	current.next = current.next.next
	return true
}

func (list *LinkedList[T]) SearchValue(i T) bool {
	if list.head == nil {
		return false
	}
	current := list.head
	for current != nil {
		if cmp.Compare(current.data, i) == 0 {
			return true
		}
		current = current.next
	}
	return false
}

func (list *LinkedList[T]) GetFirst() (T, bool) {
	var t T
	if list.head == nil {
		return t, false
	}
	return list.head.data, true
}

func (list *LinkedList[T]) GetLast() (T, bool) {
	var t T
	if list.head == nil {
		return t, false
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	return current.data, true
}

func (list *LinkedList[T]) GetSize() int {
	count := 0
	current := list.head
	for current != nil {
		count += 1
		current = current.next
	}

	return count
}

func (list *LinkedList[T]) GetItemsFromStart() []T {
	var items []T
	current := list.head
	for current != nil {
		items = append(items, current.data)
		current = current.next
	}
	return items
}

func (list *LinkedList[T]) GetItemsFromEnd() []T {
	var items []T
	current := list.tail
	for current != nil {
		items = append(items, current.data)
		current = current.prev
	}
	return items
}
