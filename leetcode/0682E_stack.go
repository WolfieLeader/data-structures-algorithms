package main

import "strconv"

// TODO:

func calPoints(operations []string) int {
	stack := make([]int, 0, len(operations))

	for _, op := range operations {
		switch op {
		case "C": // remove last
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Pop
			}

		case "D": // double last
			if len(stack) > 0 {
				last := stack[len(stack)-1]   // Peek
				stack = append(stack, 2*last) // Push
			}

		case "+": // sum of last two
			if len(stack) >= 2 {
				stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
			}

		default: // integer value
			v, err := strconv.Atoi(op)
			if err == nil {
				stack = append(stack, v)
			}
		}
	}

	sum := 0
	for _, v := range stack {
		sum += v
	}

	return sum
}
