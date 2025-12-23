package main

func fib(n int) int {
	if n < 1 {
		return 0
	}

	a, b := 0, 1
	for range n {
		a, b = b, a+b
	}

	return a
}
