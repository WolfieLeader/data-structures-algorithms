package main

// TODO:

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]

	// Phase 1 detect cycle
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]

		// Two indices are equal
		if slow == fast {
			break
		}
	}

	// Phase 2 find cycle start
	slow = nums[0] // just like setting to head of a linked list
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
