package main

// TODO:

// I: numbers=[2,7,11,15] target=9  =>  O: [1,2] (1 indexed array and not zero)
// Sorted array

func twoSum167(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return nil
}
