package main

// HACK: Core Pattern of Prefix Sums Array

// Core Pattern of Prefix Sums Array
//
// Pattern
// Precompute cumulative sums so range queries can be answered in O(1).
//
// Key idea
// Each prefix index stores the sum of all elements before it.
// prefix[i] = sum of nums[0..i-1]
//
// Why prefix has size n+1
// prefix[0] = 0 represents an empty sum.
// This removes edge cases when left = 0.
//
// Range sum formula
// sum(left..right) = prefix[right+1] - prefix[left]
//
// Why the formula works
// prefix[right+1] includes sum up to right.
// prefix[left] includes sum up to left-1.
// Subtracting removes everything before left.
//
// Time and space
// Preprocessing is O(n).
// Each query is O(1).
// Extra space is O(n).

type NumArray struct {
	prefix []int //Prefix Sums Array
}

func Constructor(nums []int) NumArray {
	prefix := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}
	return NumArray{prefix: prefix}
}

func (this *NumArray) SumRange(left int, right int) int {
	// Formula: p[right+1] - p[left] == sum(left:right)
	return this.prefix[right+1] - this.prefix[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
