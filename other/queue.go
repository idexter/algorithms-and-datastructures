package other

type Queue[T any] struct {
	data []T
}

func (q *Queue[T]) Enqueue(value T) {
	q.data = append(q.data, value)
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		var zeroVal T
		return zeroVal
	}
	firstElement := 0
	var value T = q.data[firstElement]
	q.data = q.data[1:]
	return value
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue[T]) Top() T {
	if q.IsEmpty() {
		var zeroVal T
		return zeroVal
	}
	firstElement := 0
	return q.data[firstElement]
}

type QueueOnStack[T any] struct {
	OriginalStack *Stack[T]
	InvertedStack *Stack[T]
}

func NewQueueOnStack[T any]() *QueueOnStack[T] {
	return &QueueOnStack[T]{
		OriginalStack: &Stack[T]{},
		InvertedStack: &Stack[T]{},
	}
}

func (q *QueueOnStack[T]) Enqueue(item T) {
	q.OriginalStack.Push(item)
}

func (q *QueueOnStack[T]) Dequeue() T {
	for !q.OriginalStack.IsEmpty() {
		q.InvertedStack.Push(q.OriginalStack.Pop())
	}
	top := q.InvertedStack.Pop()
	for !q.InvertedStack.IsEmpty() {
		q.OriginalStack.Push(q.InvertedStack.Pop())
	}
	return top
}

func (q *QueueOnStack[T]) IsEmpty() bool {
	return q.InvertedStack.IsEmpty() && q.OriginalStack.IsEmpty()
}

func (q *QueueOnStack[T]) Top() T {
	for !q.OriginalStack.IsEmpty() {
		q.InvertedStack.Push(q.OriginalStack.Pop())
	}
	top := q.InvertedStack.Top()
	for !q.InvertedStack.IsEmpty() {
		q.OriginalStack.Push(q.InvertedStack.Pop())
	}
	return top
}
