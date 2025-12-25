package main

// HACK: Core Pattern of Two Pointers Same Direction in Array

// Two Pointers Same Direction
//
// fast scans the array.
// slow is the write index for the next kept value.
// nums[0..slow] always holds the compressed result so far.
// For sorted arrays, keep a value only when it differs from the last kept one.

func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

		fast++
	}

	return slow + 1
}
