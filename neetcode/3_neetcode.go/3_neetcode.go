package main

func twoSum(nums []int, target int) []int {
	seenAt := make(map[int]int)
	for j, v := range nums {
		if i, found := seenAt[target-v]; found {
			return []int{i, j}
		}
		seenAt[v] = j
	}
	return []int{}
}
