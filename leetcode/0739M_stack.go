package main

func dailyTemperatures(temps []int) []int {
	answer := make([]int, len(temps))
	stack := make([]int, 0, len(temps)) // indices, temps[stack] is strictly decreasing

	for i, currTemp := range temps {
		// NOTE: LOOP until current temp is not bigger than last
		for len(stack) > 0 && currTemp > temps[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[top] = i - top
		}
		// i still needs a warmer day in the future
		stack = append(stack, i)
	}

	return answer
}
