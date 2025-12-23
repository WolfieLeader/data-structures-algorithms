package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	// a = ways to reach the previous step (i-2)
	// b = ways to reach the current step  (i-1)
	a, b := 1, 2

	// Compute ways for steps 3 through n
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}

	return b
}
