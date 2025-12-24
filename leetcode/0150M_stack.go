package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	pop := func() int {
		v := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return v
	}

	ops := map[string]func(a, b int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	for _, t := range tokens {
		if fn, ok := ops[t]; ok {
			b := pop()
			a := pop()
			stack = append(stack, fn(a, b))
		} else {
			n, _ := strconv.Atoi(t)
			stack = append(stack, n)
		}
	}

	return stack[0]
}
