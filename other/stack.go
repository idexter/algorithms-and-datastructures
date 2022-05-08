package other

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var zeroVal T
		return zeroVal
	}
	lastElement := len(s.data) - 1
	var value T = s.data[lastElement]
	s.data = s.data[:lastElement]
	return value
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Top() T {
	if s.IsEmpty() {
		var zeroVal T
		return zeroVal
	}
	lastElement := len(s.data) - 1
	return s.data[lastElement]
}
