package main

import "sort"

// TODO

func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	out := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicate first elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Early stop: if smallest is > 0, sum can't be 0
		if nums[i] > 0 {
			break
		}

		// Get the opposite of the other number
		target := -nums[i]

		left, right := i+1, len(nums)-1

		// NOTE: LOOP
		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				out = append(out, []int{nums[i], nums[left], nums[right]})

				// Move left and right to the next distinct values
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				right--
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return out
}
