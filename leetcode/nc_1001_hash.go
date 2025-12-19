package main

func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		if _, found := seen[v]; found {
			return true
		}
		seen[v] = struct{}{}
	}
	return false
}
