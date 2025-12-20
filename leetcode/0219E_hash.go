package main

// TODO:

// nums[i] == nums[j] && abs(i - j) <= k

// I: nums = [1,2,3,1],     k = 3  =>  O: true
// I: nums = [1,0,1,1],     k = 1  =>  O: true
// I: nums = [1,2,3,1,2,3], k = 2  =>  O: false

func containsNearbyDuplicate(nums []int, k int) bool {
	if k <= 0 || len(nums) <= 1 {
		return false
	}

	seenAt := make(map[int]int) // value -> last index

	for i, v := range nums {
		if j, ok := seenAt[v]; ok && i-j <= k {
			return true
		}
		seenAt[v] = i
	}
	return false
}
