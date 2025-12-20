package main

func twoSum(nums []int, target int) []int {
	seenAt := make(map[int]int)
	for i, v := range nums {
		if j, found := seenAt[target-v]; found {
			return []int{j, i}
		}
		seenAt[v] = i
	}
	return nil
}
