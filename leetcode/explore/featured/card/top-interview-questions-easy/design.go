package top_interview_questions_easy

// Problem: Min Stack.

type MinStack struct {
	data []item
}

type item struct {
	value int
	min   int
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	min := 0
	if len(s.data) > 0 {
		lastItem := s.data[len(s.data)-1]
		if val < lastItem.min {
			lastItem.min = val
			min = val
		} else {
			min = lastItem.min
		}
	} else {
		min = val
	}
	s.data = append(s.data, item{value: val, min: min})
}

func (s *MinStack) Pop() int {
	if len(s.data) == 0 {
		return 0
	}

	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]

	return last.value
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return 0
	}
	lastItem := s.data[len(s.data)-1]
	return lastItem.value
}

func (s *MinStack) GetMin() int {
	if len(s.data) == 0 {
		return 0
	}
	lastItem := s.data[len(s.data)-1]
	return lastItem.min
}

// End of Problem: Min Stack.
