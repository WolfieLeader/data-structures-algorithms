package main

// I: [1,2,4,6] => O: [48,24,12,8]
// I: [-1,0,1,2,3] => O: [0,-6,0,0]

func productExceptSelf(nums []int) []int {
	out := make([]int, len(nums))
	out[0] = 1
	// Prefix products
	for i := 1; i < len(nums); i++ {
		out[i] = out[i-1] * nums[i-1]
	}
	// Suffix products
	prod := 1 // like a carry
	for i := len(nums) - 1; i >= 0; i-- {
		out[i] *= prod
		prod *= nums[i]
	}
	return out
}
