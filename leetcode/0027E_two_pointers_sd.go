package main

// HACK: Core Pattern of Two Pointers Same Direction in Array

// Two Pointers Same Direction
//
// fast scans the array.
// slow is the write index for the next kept element.
// nums[0..slow-1] always holds the filtered result so far.
// Skip elements equal to val and overwrite in place to preserve order.

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}

		fast++
	}
	return slow
}
