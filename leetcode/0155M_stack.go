package main

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{stack: make([]int, 0), min: make([]int, 0)}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)

	if len(s.min) == 0 {
		s.min = append(s.min, val)
		return
	}

	topMin := s.min[len(s.min)-1]
	if val < topMin {
		s.min = append(s.min, val)
	} else {
		s.min = append(s.min, topMin)
	}
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.min = s.min[:len(s.min)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.min[len(s.min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
