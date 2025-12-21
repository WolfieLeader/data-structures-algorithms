package main

// Two Pointers Same Direction
//
// fast scans the array.
// slow marks the position for the next non zero element.
// nums[0..slow-1] always contains all non zero values in order.
// Swapping is safe because slow never overtakes fast.

func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			if slow != fast {
				nums[slow], nums[fast] = nums[fast], nums[slow]
			}
			slow++
		}
	}
}
