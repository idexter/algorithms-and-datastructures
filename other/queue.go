package other

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
