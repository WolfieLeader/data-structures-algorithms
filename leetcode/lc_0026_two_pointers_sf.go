package main

// TODO:

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] == nums[slow] {
			continue
		}
		slow++
		nums[slow] = nums[fast]
	}

	return slow + 1
}
