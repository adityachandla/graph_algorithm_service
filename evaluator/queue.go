package evaluator

type Queue[T any] struct {
	start, end *listNode[T]
	size       int
}

type listNode[T any] struct {
	value      T
	next, prev *listNode[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		start: nil,
		end:   nil,
		size:  0,
	}
}

func (q *Queue[T]) AddToFront(val T) {
	node := &listNode[T]{value: val, next: nil, prev: nil}
	q.size++
	if q.start != nil {
		node.next = q.start
		q.start.prev = node
	} else {
		//First insertion
		q.end = node
	}
	q.start = node
}

func (q *Queue[T]) PopBack() T {
	if q.end == nil {
		panic("Pop back called on empty queue")
	}
	q.size--
	node := q.end
	q.end = q.end.prev
	//These are important to avoid memory leak
	node.prev = nil
	node.next = nil
	return node.value
}

func (q *Queue[T]) Empty() bool {
	return q.size == 0
}
