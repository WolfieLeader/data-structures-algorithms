package main

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // value -> index

	for i, n := range nums {
		// Check if the complement (target - n) exists in the map
		if j, found := numMap[target-n]; found {
			return []int{j, i}
		}

		// Store the index of the current number
		numMap[n] = i
	}
	return nil
}
