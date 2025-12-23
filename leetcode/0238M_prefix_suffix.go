package main

// I: [1,2,4,6] => O: [48,24,12,8]
// I: [-1,0,1,2,3] => O: [0,-6,0,0]

func productExceptSelf(nums []int) []int {
	n := len(nums)
	out := make([]int, n)

	// Prefix products
	out[0] = 1
	for i := 1; i < n; i++ {
		out[i] = out[i-1] * nums[i-1]
	}

	// Suffix products
	prod := 1
	for i := n - 1; i >= 0; i-- {
		out[i] *= prod
		prod *= nums[i]
	}

	return out
}

// Another solution
func productExceptSelf2(nums []int) []int {
	n := len(nums)
	leftProd, rightProd := 1, 1
	leftArray, rightArray := make([]int, n), make([]int, n)

	for left := 0; left < n; left++ {
		right := (n - 1) - left

		leftArray[left] = leftProd
		rightArray[right] = rightProd

		leftProd *= nums[left]
		rightProd *= nums[right]
	}

	out := make([]int, n)
	for i := 0; i < n; i++ {
		out[i] = leftArray[i] * rightArray[i]
	}
	return out
}

// Imagine there is a one at the end of the array
