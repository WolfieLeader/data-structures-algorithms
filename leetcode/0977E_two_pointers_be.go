package main

// I: [-4,-1,0,3,10]  => Ex: [16,1,0,9,100]  =>  O: [0,1,9,16,100]
// I: [-7,-3,2,3,11]  => Ex: [4,9,9,49,121]  =>  O: [4,9,9,49,121]
// Sorted array => Sorted Array

func sortedSquares(nums []int) []int {
	out := make([]int, len(nums))
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[left]) > abs(nums[right]) {
			out[i] = nums[left] * nums[left]
			left++
		} else {
			out[i] = nums[right] * nums[right]
			right--
		}
	}
	return out
}
