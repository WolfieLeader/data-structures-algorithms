package main

// Core

// Greedy
//
// Keep the farthest index you can reach so far.
// Scan from left to right.
// If you ever land on an index beyond farthest, you are stuck.
//
// Update rule
// farthest = max(farthest, i + nums[i])
//
// Success
// If farthest reaches or passes the last index, return true.

func canJump(nums []int) bool {
	farthest := 0
	goal := len(nums) - 1

	for i := 0; i < len(nums); i++ {
		if i > farthest {
			return false
		}

		farthest = max(farthest, i+nums[i])
		if farthest >= goal {
			return true
		}
	}

	return true
}
