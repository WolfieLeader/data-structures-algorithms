package main

// Hash Map (Value -> Index) Lookup
//
// Iterate through the array once and store each value with its index.
// For the current value `v`, we check whether `target - v` has been seen before.
// If it has, we have found the two indices that sum to the target.
//
// The map stores values we have already passed, ensuring each pair
// uses two different indices and returns the first valid solution.
//
// Walkthrough: nums = [2,7,11,15], target = 9
// i=0, v=2  ->  need 7  ->  found=false         ->  store {2:0}
// i=1, v=7  ->  need 2  ->  found=true, index 0 ->  return [0,1]

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
