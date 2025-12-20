package main

//TODO:

// I: nums = [1,12,-5,-6,50,3], k = 4  =>  O: 12.75
// I: nums = [5],               k = 1  =>  O: 5.0

func findMaxAverage(nums []int, k int) float64 {
	// Calculate first window
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	// Sub the number to the left of the window and adding the right
	max := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if max < sum {
			max = sum
		}
	}

	return float64(max) / float64(k)
}
