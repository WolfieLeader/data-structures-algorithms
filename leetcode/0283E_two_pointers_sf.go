package main

// TODO:

// I: [0,1,0,3,12] => O: [1,3,12,0,0]

func moveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			continue
		}
		nums[slow] = nums[fast]
		slow++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}

func optimizedMoveZeroes(nums []int) {
	slow := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] == 0 {
			continue
		}
		if slow != fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
		slow++
	}
}
