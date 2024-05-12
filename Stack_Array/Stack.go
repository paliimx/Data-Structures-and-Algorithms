package StackArray

import "cmp"

const arraySize = 10

type Stack[T cmp.Ordered] struct {
	top  int
	data [arraySize]T
}

func (s *Stack[T]) Push(i T) bool {
	if s.top == len(s.data) {
		return false
	}
	s.data[s.top] = i
	s.top += 1
	return true
}

func (s *Stack[T]) Pop() (T, bool) {
	var t T
	if s.top == 0 {
		return t, false
	}
	i := s.data[s.top-1]
	s.top -= 1
	return i, true
}

func (s *Stack[T]) Peek() T {
	return s.data[s.top-1]
}

func (s *Stack[T]) Get() []T {
	return s.data[:s.top]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *Stack[T]) Empty() {
	s.top = 0
}
