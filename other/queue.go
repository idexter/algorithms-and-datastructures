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
	return q.InvertedStack.Pop()
}

func (q *QueueOnStack[T]) IsEmpty() bool {
	return q.InvertedStack.IsEmpty()
}

func (q *QueueOnStack[T]) Top() T {
	return q.InvertedStack.Top()
}
