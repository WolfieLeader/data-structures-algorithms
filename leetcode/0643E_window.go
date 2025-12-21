package main

// HACK: Core Pattern of Static Sliding Window

// Static Sliding Window
//
// Use a fixed size window of length k.
// The window moves one step at a time across the array.
//
// Key idea
// Keep a running sum of the current window.
// When the window moves:
// subtract the element leaving on the left
// add the element entering on the right
//
// Invariant
// The window always contains exactly k elements.
//
// Why it works
// Each element is added once and removed once.
// This avoids recalculating the sum for every window.
// Time complexity is linear.
//
// Walkthrough: nums = [1,12,-5,-6,50,3], k = 4
// init window [1,12,-5,-6] sum=2
// move -> [12,-5,-6,50]    sum=51
// move -> [-5,-6,50,3]     sum=42
// max sum = 51 -> average = 12.75

func findMaxAverage(nums []int, k int) float64 {
	// Calculate first window
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Sub the number to the left of the window and adding the right
	best := sum
	for right := k; right < len(nums); right++ {
		left := right - k

		sum -= nums[left]
		sum += nums[right]

		best = max(best, sum)
	}

	return float64(best) / float64(k)
}
