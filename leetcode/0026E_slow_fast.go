package main

// HACK: Core Pattern of Slow Fast Pointers in Array (aka Read Write Pointers)

// Two Pointers Same Direction
//
// fast scans the array.
// slow is the write index for the next kept value.
// nums[0..slow] always holds the compressed result so far.
// For sorted arrays, keep a value only when it differs from the last kept one.

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
	}

	return slow + 1
}
