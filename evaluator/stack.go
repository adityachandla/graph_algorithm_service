package evaluator

type Stack[T any] struct {
	arr []T
	idx int
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		arr: make([]T, 0, 10),
		idx: 0,
	}
}

func (s *Stack[T]) Empty() bool {
	return s.idx == 0
}

func (s *Stack[T]) Push(ele T) {
	if s.idx == len(s.arr) {
		s.arr = append(s.arr, ele)
	} else {
		s.arr[s.idx] = ele
	}
	s.idx++
}

func (s *Stack[T]) Pop() T {
	if s.idx == 0 {
		panic("Pop called on empty stack")
	}
	ele := s.arr[s.idx-1]
	s.idx -= 1
	return ele
}
