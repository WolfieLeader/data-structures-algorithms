package main

// Static Sliding Window with threshold check
//
// Use a fixed size window of length k.
// Slide the window across the array one step at a time.
//
// Key idea
// Track the sum of the current window.
// Compare the sum against a precomputed target.
//
// Optimization
// Instead of computing average each time,
// compare sum >= threshold * k.
//
// Invariant
// The window always contains exactly k elements.
//
// Why it works
// Each element enters and leaves the window once.
// All checks are constant time.
// Overall time complexity is linear.
//
// Walkthrough: arr = [2,2,2,2,5,5,5,8], k=3, threshold=4
// target = 12
// window [2,2,2] sum=6   < target
// window [2,2,2] sum=6   < target
// window [2,2,5] sum=9   < target
// window [2,5,5] sum=12  >= target count++
// window [5,5,5] sum=15  >= target count++
// window [5,5,8] sum=18  >= target count++
// result = 3

func numOfSubarrays(arr []int, k int, threshold int) int {
	target := threshold * k

	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	count := 0
	if sum >= target {
		count++
	}

	for right := k; right < len(arr); right++ {
		left := right - k

		sum -= arr[left]
		sum += arr[right]

		if sum >= target {
			count++
		}
	}
	return count
}
